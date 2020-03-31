// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package client

import (
	"io"

	"github.com/mattermost/mattermost-server/v5/model"
)

type Client interface {
	CreateChannel(channel *model.Channel) (*model.Channel, *model.Response)
	RemoveUserFromChannel(channelId, userId string) (bool, *model.Response)
	GetChannelMembers(channelId string, page, perPage int, etag string) (*model.ChannelMembers, *model.Response)
	AddChannelMember(channelId, userId string) (*model.ChannelMember, *model.Response)
	DeleteChannel(channelId string) (bool, *model.Response)
	GetPublicChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response)
	GetDeletedChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response)
	RestoreChannel(channelId string) (*model.Channel, *model.Response)
	ConvertChannelToPrivate(channelId string) (*model.Channel, *model.Response)
	PatchChannel(channelId string, patch *model.ChannelPatch) (*model.Channel, *model.Response)
	GetChannelByName(channelName, teamId string, etag string) (*model.Channel, *model.Response)
	GetChannelByNameIncludeDeleted(channelName, teamId string, etag string) (*model.Channel, *model.Response)
	GetChannel(channelId, etag string) (*model.Channel, *model.Response)
	GetTeam(teamId, etag string) (*model.Team, *model.Response)
	GetTeamByName(name, etag string) (*model.Team, *model.Response)
	GetAllTeams(etag string, page int, perPage int) ([]*model.Team, *model.Response)
	CreateTeam(team *model.Team) (*model.Team, *model.Response)
	PatchTeam(teamId string, patch *model.TeamPatch) (*model.Team, *model.Response)
	AddTeamMember(teamId, userId string) (*model.TeamMember, *model.Response)
	RemoveTeamMember(teamId, userId string) (bool, *model.Response)
	SoftDeleteTeam(teamId string) (bool, *model.Response)
	PermanentDeleteTeam(teamId string) (bool, *model.Response)
	SearchTeams(search *model.TeamSearch) ([]*model.Team, *model.Response)
	GetPost(postId string, etag string) (*model.Post, *model.Response)
	CreatePost(post *model.Post) (*model.Post, *model.Response)
	GetPostsForChannel(channelId string, page, perPage int, etag string) (*model.PostList, *model.Response)
	GetLdapGroups() ([]*model.Group, *model.Response)
	GetGroupsByChannel(channelId string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response)
	GetGroupsByTeam(teamId string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response)
	UploadLicenseFile(data []byte) (bool, *model.Response)
	RemoveLicenseFile() (bool, *model.Response)
	GetLogs(page, perPage int) ([]string, *model.Response)
	GetRoleByName(name string) (*model.Role, *model.Response)
	PatchRole(roleId string, patch *model.RolePatch) (*model.Role, *model.Response)
	UploadPlugin(file io.Reader) (*model.Manifest, *model.Response)
	RemovePlugin(id string) (bool, *model.Response)
	EnablePlugin(id string) (bool, *model.Response)
	DisablePlugin(id string) (bool, *model.Response)
	GetPlugins() (*model.PluginsResponse, *model.Response)
	GetUser(userId, etag string) (*model.User, *model.Response)
	GetUserByUsername(userName, etag string) (*model.User, *model.Response)
	GetUserByEmail(email, etag string) (*model.User, *model.Response)
	CreateUser(user *model.User) (*model.User, *model.Response)
	UpdateUserRoles(userId, roles string) (bool, *model.Response)
	InviteUsersToTeam(teamId string, userEmails []string) (bool, *model.Response)
	SendPasswordResetEmail(email string) (bool, *model.Response)
	UpdateUser(user *model.User) (*model.User, *model.Response)
	UpdateUserMfa(userId, code string, activate bool) (bool, *model.Response)
	CreateUserAccessToken(userId, description string) (*model.UserAccessToken, *model.Response)
	RevokeUserAccessToken(tokenId string) (bool, *model.Response)
	GetUserAccessTokensForUser(userId string, page, perPage int) ([]*model.UserAccessToken, *model.Response)
	CreateCommand(cmd *model.Command) (*model.Command, *model.Response)
	ListCommands(teamId string, customOnly bool) ([]*model.Command, *model.Response)
	GetCommandById(cmdId string) (*model.Command, *model.Response)
	UpdateCommand(cmd *model.Command) (*model.Command, *model.Response)
	MoveCommand(teamId string, commandId string) (bool, *model.Response)
	DeleteCommand(commandId string) (bool, *model.Response)
	GetConfig() (*model.Config, *model.Response)
	UpdateConfig(*model.Config) (*model.Config, *model.Response)
	SyncLdap() (bool, *model.Response)
	GetUsers(page, perPage int, etag string) ([]*model.User, *model.Response)
	GetUsersByIds(userIds []string) ([]*model.User, *model.Response)
	UpdateUserActive(userId string, activate bool) (bool, *model.Response)
	UpdateTeam(team *model.Team) (*model.Team, *model.Response)
	UpdateChannelPrivacy(channelId string, privacy string) (*model.Channel, *model.Response)
	CreateBot(bot *model.Bot) (*model.Bot, *model.Response)
	PatchBot(userId string, patch *model.BotPatch) (*model.Bot, *model.Response)
	GetBots(page, perPage int, etag string) ([]*model.Bot, *model.Response)
	GetBotsIncludeDeleted(page, perPage int, etag string) ([]*model.Bot, *model.Response)
	GetBotsOrphaned(page, perPage int, etag string) ([]*model.Bot, *model.Response)
	DisableBot(botUserId string) (*model.Bot, *model.Response)
	EnableBot(botUserId string) (*model.Bot, *model.Response)
	AssignBot(botUserId, newOwnerId string) (*model.Bot, *model.Response)
	SetServerBusy(secs int) (bool, *model.Response)
	ClearServerBusy() (bool, *model.Response)
	GetServerBusy() (*model.ServerBusyState, *model.Response)
}
