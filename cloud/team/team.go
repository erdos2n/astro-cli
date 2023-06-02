package team

import (
	httpContext "context"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"

	astrocore "github.com/astronomer/astro-cli/astro-client-core"
	"github.com/astronomer/astro-cli/cloud/user"
	"github.com/astronomer/astro-cli/context"
	"github.com/astronomer/astro-cli/pkg/input"
	"github.com/astronomer/astro-cli/pkg/printutil"
)

var (
	errInvalidTeamKey    = errors.New("invalid team selection")
	ErrInvalidName       = errors.New("no name provided for the team. Retry with a valid name")
	ErrTeamNotFound      = errors.New("no team was found for the ID you provided")
	ErrWrongEnforceInput = errors.New("the input to the `--enforce-cicd` flag")
	ErrNoShortName       = errors.New("cannot retrieve organization short name from context")
	teamPagnationLimit   = 100
)

func newTableOut() *printutil.Table {
	return &printutil.Table{
		Padding:        []int{44, 50},
		DynamicPadding: true,
		Header:         []string{"NAME", "ID"},
		ColorRowCode:   [2]string{"\033[1;32m", "\033[0m"},
	}
}

func CreateTeam(name string, description string, out io.Writer, client astrocore.CoreClient) error {
	if name == "" {
		return ErrInvalidName
	}
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if ctx.OrganizationShortName == "" {
		return user.ErrNoShortName
	}
	teamCreateRequest := astrocore.CreateTeamJSONRequestBody{
		Description: &description,
		Name:        name,
	}
	resp, err := client.CreateTeamWithResponse(httpContext.Background(), ctx.OrganizationShortName, teamCreateRequest)
	if err != nil {
		return err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Astro Team %s was successfully created\n", name)
	return nil
}

func UpdateWorkspaceTeamRole(id, role, workspace string, out io.Writer, client astrocore.CoreClient) error {
	err := user.IsWorkspaceRoleValid(role)
	if err != nil {
		return err
	}
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if ctx.OrganizationShortName == "" {
		return user.ErrNoShortName
	}
	if workspace == "" {
		workspace = ctx.Workspace
	}

	teams, err := GetWorkspaceTeams(client, workspace, teamPagnationLimit)

	if err != nil {
		return err
	}
	var team astrocore.Team
	if id == "" {
		team, err = selectTeam(teams)
		if err != nil {
			return err
		}
	} else {
		for i := range teams {
			if teams[i].Id == id {
				team = teams[i]
			}
		}

		if team.Id == "" {
			return ErrTeamNotFound
		}
	}
	teamID := team.Id

	teamMutateRequest := astrocore.MutateWorkspaceTeamRoleRequest{Role: role}
	resp, err := client.MutateWorkspaceTeamRoleWithResponse(httpContext.Background(), ctx.OrganizationShortName, workspace, teamID, teamMutateRequest)

	if err != nil {
		return err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "The workspace team %s role was successfully updated to %s\n", teamID, role)
	return nil
}

func UpdateTeam(id, name, description string, out io.Writer, client astrocore.CoreClient) error {
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if ctx.OrganizationShortName == "" {
		return user.ErrNoShortName
	}
	teams, err := GetOrgTeams(client)
	if err != nil {
		return err
	}
	var team astrocore.Team
	if id == "" {
		team, err = selectTeam(teams)
		if err != nil {
			return err
		}
	} else {
		for i := range teams {
			if teams[i].Id == id {
				team = teams[i]
			}
		}
		if team.Id == "" {
			return ErrTeamNotFound
		}
	}
	teamID := team.Id

	teamUpdateRequest := astrocore.UpdateTeamJSONRequestBody{}

	if name == "" {
		teamUpdateRequest.Name = team.Name
	} else {
		teamUpdateRequest.Name = name
	}

	if description == "" {
		teamUpdateRequest.Description = *team.Description
	} else {
		teamUpdateRequest.Description = description
	}

	resp, err := client.UpdateTeamWithResponse(httpContext.Background(), ctx.OrganizationShortName, teamID, teamUpdateRequest)
	if err != nil {
		return err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Astro Team %s was successfully updated\n", team.Name)
	return nil
}

func RemoveWorkspaceTeam(id, workspace string, out io.Writer, client astrocore.CoreClient) error {
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if ctx.OrganizationShortName == "" {
		return user.ErrNoShortName
	}
	if workspace == "" {
		workspace = ctx.Workspace
	}
	teams, err := GetWorkspaceTeams(client, workspace, teamPagnationLimit)
	if err != nil {
		return err
	}
	var team astrocore.Team
	if id == "" {
		team, err = selectTeam(teams)
		if err != nil {
			return err
		}
	} else {
		for i := range teams {
			if teams[i].Id == id {
				team = teams[i]
			}
		}
		if team.Id == "" {
			return ErrTeamNotFound
		}
	}
	teamID := team.Id
	resp, err := client.DeleteWorkspaceTeamWithResponse(httpContext.Background(), ctx.OrganizationShortName, ctx.Workspace, teamID)
	if err != nil {
		return err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Astro Team %s was successfully removed from workspace %s\n", team.Name, ctx.Workspace)
	return nil
}

func selectTeam(teams []astrocore.Team) (astrocore.Team, error) {
	table := printutil.Table{
		Padding:        []int{30, 50, 10, 50, 10, 10, 10},
		DynamicPadding: true,
		Header:         []string{"#", "TEAMNAME", "ID"},
	}

	fmt.Println("\nPlease select the team you would like to update:")

	teamMap := map[string]astrocore.Team{}
	for i := range teams {
		index := i + 1
		table.AddRow([]string{
			strconv.Itoa(index),
			teams[i].Name,
			teams[i].Id,
		}, false)
		teamMap[strconv.Itoa(index)] = teams[i]
	}

	table.Print(os.Stdout)
	choice := input.Text("\n> ")
	selected, ok := teamMap[choice]
	if !ok {
		return astrocore.Team{}, errInvalidTeamKey
	}
	return selected, nil
}

func GetWorkspaceTeams(client astrocore.CoreClient, workspace string, limit int) ([]astrocore.Team, error) {
	offset := 0
	var teams []astrocore.Team

	ctx, err := context.GetCurrentContext()
	if err != nil {
		return nil, err
	}
	if ctx.OrganizationShortName == "" {
		return nil, ErrNoShortName
	}
	if workspace == "" {
		workspace = ctx.Workspace
	}
	for {
		resp, err := client.ListWorkspaceTeamsWithResponse(httpContext.Background(), ctx.OrganizationShortName, workspace, &astrocore.ListWorkspaceTeamsParams{
			Offset: &offset,
			Limit:  &limit,
		})
		if err != nil {
			return nil, err
		}
		err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
		if err != nil {
			return nil, err
		}
		teams = append(teams, resp.JSON200.Teams...)

		if resp.JSON200.TotalCount <= offset {
			break
		}

		offset += limit
	}

	return teams, nil
}

// Prints a list of all of an organizations users
func ListWorkspaceTeams(out io.Writer, client astrocore.CoreClient, workspace string) error {
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if workspace == "" {
		workspace = ctx.Workspace
	}
	table := printutil.Table{
		Padding:        []int{10, 50, 50, 10, 10, 10, 10},
		DynamicPadding: true,
		Header:         []string{"ID", "Role", "Name", "Description", "CREATE DATE"},
	}
	teams, err := GetWorkspaceTeams(client, workspace, teamPagnationLimit)
	if err != nil {
		return err
	}

	for i := range teams {
		var teamRole string
		for _, role := range *teams[i].Roles {
			if role.EntityType == "WORKSPACE" && role.EntityId == workspace {
				teamRole = role.Role
			}
		}
		table.AddRow([]string{
			teams[i].Id,
			teamRole,
			teams[i].Name,
			*teams[i].Description,
			teams[i].CreatedAt.Format(time.RFC3339),
		}, false)
	}

	table.Print(out)
	return nil
}

func AddWorkspaceTeam(id, role, workspace string, out io.Writer, client astrocore.CoreClient) error {
	err := user.IsWorkspaceRoleValid(role)
	if err != nil {
		return err
	}
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if ctx.OrganizationShortName == "" {
		return ErrNoShortName
	}
	if workspace == "" {
		workspace = ctx.Workspace
	}

	// Get all org team. Setting limit to 1000 for now
	teams, err := GetOrgTeams(client)
	if err != nil {
		return err
	}
	teamID, err := getTeamID(id, teams)
	if err != nil {
		return err
	}

	mutateUserInput := astrocore.MutateWorkspaceTeamRoleRequest{
		Role: role,
	}
	resp, err := client.MutateWorkspaceTeamRoleWithResponse(httpContext.Background(), ctx.OrganizationShortName, workspace, teamID, mutateUserInput)
	if err != nil {
		return err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "The team %s was successfully added to the workspace with the role %s\n", teamID, role)
	return nil
}

// Returns a list of all of an organizations teams
func GetOrgTeams(client astrocore.CoreClient) ([]astrocore.Team, error) {
	offset := 0
	var teams []astrocore.Team
	ctx, err := context.GetCurrentContext()

	if err != nil {
		return nil, err
	}
	if ctx.OrganizationShortName == "" {
		return nil, ErrNoShortName
	}

	for {
		resp, err := client.ListOrganizationTeamsWithResponse(httpContext.Background(), ctx.OrganizationShortName, &astrocore.ListOrganizationTeamsParams{
			Offset: &offset,
			Limit:  &teamPagnationLimit,
		})
		if err != nil {
			return nil, err
		}
		err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
		if err != nil {
			return nil, err
		}
		teams = append(teams, resp.JSON200.Teams...)

		if resp.JSON200.TotalCount <= offset {
			break
		}

		offset += teamPagnationLimit
	}

	return teams, nil
}

// Prints a list of all of an organizations users
func ListOrgTeams(out io.Writer, client astrocore.CoreClient) error {
	table := printutil.Table{
		Padding:        []int{10, 50, 50, 10, 10, 10},
		DynamicPadding: true,
		Header:         []string{"ID", "Name", "Description", "CREATE DATE"},
	}
	teams, err := GetOrgTeams(client)

	if err != nil {
		return err
	}

	for i := range teams {
		table.AddRow([]string{
			teams[i].Id,
			teams[i].Name,
			*teams[i].Description,
			teams[i].CreatedAt.Format(time.RFC3339),
		}, false)
	}

	table.Print(out)
	return nil
}

func getTeamID(id string, teams []astrocore.Team) (teamID string, err error) {
	if id == "" {
		team, err := selectTeam(teams)
		teamID = team.Id
		if err != nil {
			return teamID, err
		}
	} else {
		for i := range teams {
			if teams[i].Id == id {
				teamID = teams[i].Id
			}
		}
	}
	return teamID, nil
}

func Delete(id string, out io.Writer, client astrocore.CoreClient) error {
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}
	if ctx.OrganizationShortName == "" {
		return user.ErrNoShortName
	}
	teams, err := GetOrgTeams(client)
	if err != nil {
		return err
	}
	var team astrocore.Team
	if id == "" {
		team, err = selectTeam(teams)
		if err != nil {
			return err
		}
	} else {
		for i := range teams {
			if teams[i].Id == id {
				team = teams[i]
			}
		}
		if team.Id == "" {
			return ErrTeamNotFound
		}
	}
	teamID := team.Id
	resp, err := client.DeleteTeamWithResponse(httpContext.Background(), ctx.OrganizationShortName, teamID)
	if err != nil {
		return err
	}
	err = astrocore.NormalizeAPIError(resp.HTTPResponse, resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Astro Team %s was successfully deleted\n", team.Name)
	return nil
}
