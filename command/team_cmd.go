package command

import (
	"context"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
	gcli "github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)


func NewTeamClient(c *gcli.Context) *team.Client {
	teamCli, cliErr := team.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the team client. " + cliErr.Error()
		fmt.Printf("%s\n", message)
		os.Exit(1)
	}
	printVerboseMessage("Team Client created.")
	return teamCli
}

func CreateTeamAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	createTeamRequest := &team.CreateTeamRequest{}

	if teamName, ok := getVal("name",c );ok {
		createTeamRequest.Name = teamName
	}
	if desc, ok := getVal("desc", c); ok {
		createTeamRequest.Description = desc
	}

	userName, _ := getVal("userName",c)
	userID, _ := getVal("userId",c)

	createTeamRequest.Members = []team.Member{
		{
			User: team.User{
				Username: userName,
				ID: userID,
			},
		},
	}
	if role, ok := getVal("role", c);ok {
		createTeamRequest.Members[0].Role = role
	}

	printResponse(createTeamRequest, nil, c)

	resp, err := teamCli.Create(context.Background(), createTeamRequest)
	printResponse(resp, err, c)
}

func UpdateTeamAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	updateTeamRequest := &team.UpdateTeamRequest{}

	if teamName, ok := getVal("name",c );ok {
		updateTeamRequest.Name = teamName
	}
	if desc, ok := getVal("desc", c); ok {
		updateTeamRequest.Description = desc
	}
	if ID, ok := getVal("id", c); ok {
		updateTeamRequest.Id = ID
	}

	userName, _ := getVal("userName",c)
	userID, _ := getVal("userId",c)
	updateTeamRequest.Members = []team.Member{
		{
			User: team.User{
				Username: userName,
				ID: userID,
			},
		},
	}
	if role, ok := getVal("role", c);ok {
		updateTeamRequest.Members[0].Role = role
	}

	resp, err := teamCli.Update(context.Background(), updateTeamRequest)
	printResponse(resp, err, c)
}


func GetTeamAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	getTeamRequest := &team.GetTeamRequest{}

	if teamName, ok := getVal("name", c); ok {
		getTeamRequest = &team.GetTeamRequest{
			IdentifierType: team.Name,
			IdentifierValue: teamName,
		}
	} else if teamID, ok:= getVal("id", c); ok{
		getTeamRequest = &team.GetTeamRequest{
			IdentifierType: team.Id,
			IdentifierValue: teamID,
		}
	}

	resp, err := teamCli.Get(context.Background(), getTeamRequest)
	printResponse(resp, err, c)
}

func DeleteTeamAction(c *gcli.Context){
	teamCli := NewTeamClient(c)
	var deleteTeamRequest *team.DeleteTeamRequest

	if teamName, ok := getVal("name", c); ok {
		deleteTeamRequest = &team.DeleteTeamRequest{
			IdentifierType: team.Name,
			IdentifierValue: teamName,
		}
	} else if teamID, ok:= getVal("id", c); ok{
		deleteTeamRequest = &team.DeleteTeamRequest{
			IdentifierType: team.Id,
			IdentifierValue: teamID,
		}
	}
	resp, err := teamCli.Delete(context.Background(), deleteTeamRequest)
	printResponse(resp, err, c)
}

func ListTeamsAction(c *gcli.Context){
	teamCli := NewTeamClient(c)
	resp, err := teamCli.List(context.Background(), &team.ListTeamRequest{})
	printResponse(resp, err, c)
}

func ListRolesAction(c *gcli.Context)  {
	teamCli := NewTeamClient(c)
	listTeamRolesRequest := &team.ListTeamRoleRequest{}

	if teamName, ok := getVal("name",c); ok {
		listTeamRolesRequest = &team.ListTeamRoleRequest{
			TeamIdentifierType: team.Name,
			TeamIdentifierValue: teamName,
		}
	} else if teamID, ok := getVal("id", c); ok {
		listTeamRolesRequest = &team.ListTeamRoleRequest{
			TeamIdentifierType: team.Id,
			TeamIdentifierValue: teamID,
		}
	}

	resp, err := teamCli.ListRole(context.Background(), listTeamRolesRequest)
	printResponse(resp, err, c)
}

func CreateRoleAction(c *gcli.Context)  {
	teamCli := NewTeamClient(c)
	createTeamRoleRequest := &team.CreateTeamRoleRequest{}

	if teamName, ok := getVal("name",c); ok {
		createTeamRoleRequest.TeamIdentifierType = team.Name
		createTeamRoleRequest.TeamIdentifierValue = teamName
	} else if teamID, ok := getVal("id", c); ok {
		createTeamRoleRequest.TeamIdentifierType = team.Id
		createTeamRoleRequest.TeamIdentifierValue =  teamID
	}
	if roleName, ok := getVal("roleName",c); ok {
		createTeamRoleRequest.Name = roleName
	}

	roleRights := []team.Right{}
	granted := true
	if rightsArgVal, ok := getVal("rights", c); ok {
		rights := strings.Split(rightsArgVal, ",")
		for _, right := range rights{
			roleRights = append(roleRights, team.Right{
				Right:   right,
				Granted: &granted,
			})
		}
	}
	createTeamRoleRequest.Rights = roleRights

	resp, err := teamCli.CreateRole(context.Background(), createTeamRoleRequest)
	printResponse(resp, err, c)
}

func ListRoleRightsAction(c *gcli.Context){
	type right struct {
		Name string `json:"name"`
		Description string `json:"description"`
		Category string `json:"category"`
	}
	roleRights := []right{
		{
			Name: "manage-members",
			Description: "Manage Team Members",
			Category: "Member Management",
		},
		{
			Name: "edit-team-roles",
			Description: "reate/Update Team Roles",
			Category: "Member Management",
		},
		{
			Name: "delete-team-roles",
			Description: "Delete Team Roles",
			Category: "Member Management",
		},
		{
			Name: "access-member-profiles",
			Description: "Access Profiles of Team Members",
			Category: "Member Management",
		},
		{
			Name: "edit-member-profiles",
			Description: "Edit Profiles of Team Members",
			Category: "Member Management",
		},
		{
			Name: "edit-routing-rules",
			Description: "Create/Update Routing Rules",
			Category: "Configurations",
		},
		{
			Name: "delete-routing-rules",
			Description: "Delete Routing Rules",
			Category: "Configurations",
		},
		{
			Name: "edit-escalations",
			Description: "Create/Update Escalations",
			Category: "Configurations",
		},
		{
			Name: "delete-escalations",
			Description: "Delete Escalations",
			Category: "Configurations",
		},
		{
			Name: "edit-schedules",
			Description: "Create/Update Schedules",
			Category: "Configurations",
		},
		{
			Name: "delete-schedules",
			Description: "Delete Schedules",
			Category: "Configurations",
		},
		{
			Name: "edit-integrations",
			Description: "Create/Update Integrations",
			Category: "Configurations",
		},
		{
			Name: "delete-integrations",
			Description: "Delete Integrations",
			Category: "Configurations",
		},
		{
			Name: "edit-automation-actions",
			Description: "Create/Update Automation Actions",
			Category: "Configurations",
		},
		{
			Name: "delete-automation-actions",
			Description: "Delete Automation Actions",
			Category: "Configurations",
		},
		{
			Name: "edit-heartbeats",
			Description: "Create/Update Heartbeats",
			Category: "Configurations",
		},
		{
			Name: "delete-heartbeats",
			Description: "Delete Heartbeats",
			Category: "Configurations",
		},
		{
			Name: "edit-policies",
			Description: "Create/Update Policies",
			Category: "Configurations",
		},
		{
			Name: "delete-policies",
			Description: "Delete Policies",
			Category: "Configurations",
		},
		{
			Name: "edit-maintenance",
			Description: "Create/Update Maintenance",
			Category: "Configurations",
		},
		{
			Name: "delete-maintenance",
			Description: "Delete Maintenance",
			Category: "Configurations",
		},
		{
			Name: "access-reports",
			Description: "Access Reports",
			Category: "Configurations",
		},
		{
			Name: "edit-services",
			Description: "Create/Update Services",
			Category: "Incident Configurations",
		},
		{
			Name: "delete-services",
			Description: "Delete Services",
			Category: "Incident Configurations",
		},
		{
			Name: "edit-rooms",
			Description: "Create/Update Rooms",
			Category: "Incident Configurations",
		},
		{
			Name: "delete-rooms",
			Description: "Delete Rooms",
			Category: "Incident Configurations",
		},
		{
			Name: "subscription-to-services",
			Description: "Subscription To Services",
			Category: "Incident Configurations",
		},
	}
	printResponse(struct {
		Rights []right `json:"rights"`
	}{roleRights}, nil , c)
}

func ListTeamRoutingRulesAction(c *gcli.Context)  {
	teamCli:= NewTeamClient(c)
	listRoutingRulesRequest := &team.ListRoutingRulesRequest{}

	if teamName, ok := getVal("name",c); ok {
		listRoutingRulesRequest = &team.ListRoutingRulesRequest{
			TeamIdentifierType: team.Name,
			TeamIdentifierValue: teamName,
		}
	} else if teamID, ok := getVal("id", c); ok {
		listRoutingRulesRequest = &team.ListRoutingRulesRequest{
			TeamIdentifierType: team.Id,
			TeamIdentifierValue: teamID,
		}
	}

	resp, err := teamCli.ListRoutingRules(context.Background(), listRoutingRulesRequest)
	printResponse(resp, err, c)
}

func DeleteTeamRoutingRuleAction(c *gcli.Context)  {
	teamCli := NewTeamClient(c)
	deleteRoutingRuleRequest := &team.DeleteRoutingRuleRequest{}

	if teamName, ok := getVal("name",c); ok {
		deleteRoutingRuleRequest = &team.DeleteRoutingRuleRequest{
			TeamIdentifierType: team.Name,
			TeamIdentifierValue: teamName,
		}
	} else if teamID, ok := getVal("id", c); ok {
		deleteRoutingRuleRequest = &team.DeleteRoutingRuleRequest{
			TeamIdentifierType: team.Id,
			TeamIdentifierValue: teamID,
		}
	}
	if ruleID, ok := getVal("ruleId", c); ok {
		deleteRoutingRuleRequest.RoutingRuleId = ruleID
	}

	resp, err := teamCli.DeleteRoutingRule(context.Background(), deleteRoutingRuleRequest)
	printResponse(resp, err, c)
}


func GetTeamRoleAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	getTeamRoleRequest := &team.GetTeamRoleRequest{}

	if teamName, ok := getVal("teamName",c); ok {
		getTeamRoleRequest.TeamName = teamName
	} else if teamID, ok := getVal("teamId",c);ok {
		getTeamRoleRequest.TeamID = teamID
	}

	if roleName, ok := getVal("roleName",c); ok {
		getTeamRoleRequest.RoleName = roleName
	} else if roleID, ok := getVal("roleId",c);ok {
		getTeamRoleRequest.RoleID = roleID
	}

	resp, err := teamCli.GetRole(context.Background(), getTeamRoleRequest)
	printResponse(resp, err, c)
}


func DeleteTeamRoleAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	deleteTeamRoleRequest := &team.DeleteTeamRoleRequest{}

	if teamName, ok := getVal("teamName",c); ok {
		deleteTeamRoleRequest.TeamName = teamName
	} else if teamID, ok := getVal("teamId",c);ok {
		deleteTeamRoleRequest.TeamID = teamID
	}

	if roleName, ok := getVal("roleName",c); ok {
		deleteTeamRoleRequest.RoleName = roleName
	} else if roleID, ok := getVal("roleId",c);ok {
		deleteTeamRoleRequest.RoleID = roleID
	}

	resp, err := teamCli.DeleteRole(context.Background(), deleteTeamRoleRequest)
	printResponse(resp, err, c)
}

func AddMemberAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	addMemberRequest := &team.AddTeamMemberRequest{}

	if teamName, ok := getVal("teamName",c); ok {
		addMemberRequest.TeamIdentifierType = team.Name
		addMemberRequest.TeamIdentifierValue = teamName
	} else if teamID, ok := getVal("teamId", c); ok {
		addMemberRequest.TeamIdentifierType = team.Id
		addMemberRequest.TeamIdentifierValue = teamID
	}

	if role, ok := getVal("role",c); ok {
		addMemberRequest.Role = role
	}

	if userID, ok := getVal("userId",c); ok {
		addMemberRequest.User.ID = userID
	} else if userName, ok := getVal("userName",c); ok {
		addMemberRequest.User.Username = userName
	}

	resp, err := teamCli.AddMember(context.Background(), addMemberRequest)
	printResponse(resp, err, c)
}

func RemoveMemberAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	removeMemberRequest := &team.RemoveTeamMemberRequest{}

	if teamName, ok := getVal("teamName",c); ok {
		removeMemberRequest.TeamIdentifierType = team.Name
		removeMemberRequest.TeamIdentifierValue = teamName
	} else if teamID, ok := getVal("teamId", c); ok {
		removeMemberRequest.TeamIdentifierType = team.Id
		removeMemberRequest.TeamIdentifierValue = teamID
	}

	if userID, ok := getVal("userId",c); ok {
		removeMemberRequest.MemberIdentifierType = team.Id
		removeMemberRequest.MemberIdentifierValue = userID
	} else if userName, ok := getVal("userName",c); ok {
		removeMemberRequest.MemberIdentifierType = team.Username
		removeMemberRequest.MemberIdentifierValue = userName
	}

	resp, err := teamCli.RemoveMember(context.Background(), removeMemberRequest)
	printResponse(resp, err, c)
}

func GetRoutingRuleAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	getRoutingRuleRequest := &team.GetRoutingRuleRequest{}

	if teamName, ok := getVal("teamName",c); ok {
		getRoutingRuleRequest.TeamIdentifierType = team.Name
		getRoutingRuleRequest.TeamIdentifierValue = teamName
	} else if teamID, ok := getVal("teamId", c); ok {
		getRoutingRuleRequest.TeamIdentifierType = team.Id
		getRoutingRuleRequest.TeamIdentifierValue = teamID
	}

	if ruleID, ok := getVal("ruleId",c); ok {
		getRoutingRuleRequest.RoutingRuleId = ruleID
	}

	resp, err := teamCli.GetRoutingRule(context.Background(), getRoutingRuleRequest)
	printResponse(resp, err, c)
}

func ListTeamLogsAction(c *gcli.Context) {
	teamCli := NewTeamClient(c)
	listLogsRequest := &team.ListTeamLogsRequest{}

	if teamName, ok := getVal("name", c); ok {
			listLogsRequest.IdentifierType = team.Name
			listLogsRequest.IdentifierValue = teamName
	} else if teamID, ok:= getVal("id", c); ok{
		listLogsRequest.IdentifierType = team.Id
		listLogsRequest.IdentifierValue = teamID
	}

	if limit, ok := getVal("limit", c); ok {
		limit, err := strconv.Atoi(limit)
		if err != nil {
			os.Exit(2)
		}
		listLogsRequest.Limit = limit
	}
	if offset, ok := getVal("offset", c); ok {
		offset, err := strconv.Atoi(offset)
		if err != nil {
			os.Exit(2)
		}
		listLogsRequest.Offset = offset
	}
	if order, ok := getVal("order", c); ok {
		listLogsRequest.Order = order
	}
	resp, err := teamCli.ListTeamLogs(context.Background(), listLogsRequest)
	printResponse(resp, err, c)
}

func printResponse(resp interface{},err error, c *gcli.Context) {
	if err != nil {
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(output)
}
