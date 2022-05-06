package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gitopia/gitopia/x/gitopia/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagExpiration             = "expiration"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateTask())
	cmd.AddCommand(CmdUpdateTask())
	cmd.AddCommand(CmdDeleteTask())
	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdCreateRelease())
	cmd.AddCommand(CmdUpdateRelease())
	cmd.AddCommand(CmdDeleteRelease())

	cmd.AddCommand(CmdCreatePullRequest())
	cmd.AddCommand(CmdUpdatePullRequest())
	cmd.AddCommand(CmdUpdatePullRequestTitle())
	cmd.AddCommand(CmdUpdatePullRequestDescription())
	cmd.AddCommand(CmdInvokeMergePullRequest())
	cmd.AddCommand(CmdSetPullRequestState())
	cmd.AddCommand(CmdAddPullRequestAssignees())
	cmd.AddCommand(CmdRemovePullRequestAssignees())
	cmd.AddCommand(CmdAddPullRequestReviewers())
	cmd.AddCommand(CmdRemovePullRequestReviewers())
	cmd.AddCommand(CmdAddPullRequestLabels())
	cmd.AddCommand(CmdRemovePullRequestLabels())
	cmd.AddCommand(CmdDeletePullRequest())

	cmd.AddCommand(CmdCreateOrganization())
	cmd.AddCommand(CmdRenameOrganization())
	cmd.AddCommand(CmdUpdateOrganizationMember())
	cmd.AddCommand(CmdRemoveOrganizationMember())
	cmd.AddCommand(CmdUpdateOrganization())
	cmd.AddCommand(CmdUpdateOrganizationDescription())
	cmd.AddCommand(CmdDeleteOrganization())

	cmd.AddCommand(CmdCreateComment())
	cmd.AddCommand(CmdUpdateComment())
	cmd.AddCommand(CmdDeleteComment())

	cmd.AddCommand(CmdCreateIssue())
	cmd.AddCommand(CmdUpdateIssue())
	cmd.AddCommand(CmdUpdateIssueTitle())
	cmd.AddCommand(CmdUpdateIssueDescription())
	cmd.AddCommand(CmdToggleIssueState())
	cmd.AddCommand(CmdAddIssueAssignees())
	cmd.AddCommand(CmdRemoveIssueAssignees())
	cmd.AddCommand(CmdAddIssueLabels())
	cmd.AddCommand(CmdRemoveIssueLabels())
	cmd.AddCommand(CmdDeleteIssue())

	cmd.AddCommand(CmdCreateRepository())
	cmd.AddCommand(CmdInvokeForkRepository())
	cmd.AddCommand(CmdAuthorizeGitServer())
	cmd.AddCommand(CmdForkRepository())
	cmd.AddCommand(CmdRenameRepository())
	cmd.AddCommand(CmdChangeOwner())
	cmd.AddCommand(CmdUpdateRepositoryCollaborator())
	cmd.AddCommand(CmdRemoveRepositoryCollaborator())
	cmd.AddCommand(CmdCreateRepositoryLabel())
	cmd.AddCommand(CmdUpdateRepositoryLabel())
	cmd.AddCommand(CmdDeleteRepositoryLabel())
	cmd.AddCommand(CmdSetRepositoryBranch())
	cmd.AddCommand(CmdSetDefaultBranch())
	cmd.AddCommand(CmdDeleteBranch())
	cmd.AddCommand(CmdSetRepositoryTag())
	cmd.AddCommand(CmdDeleteTag())
	cmd.AddCommand(CmdToggleRepositoryForking())
	cmd.AddCommand(CmdDeleteRepository())

	cmd.AddCommand(CmdCreateUser())
	cmd.AddCommand(CmdUpdateUser())
	cmd.AddCommand(CmdUpdateUserBio())
	cmd.AddCommand(CmdUpdateUserAvatar())
	cmd.AddCommand(CmdDeleteUser())
	cmd.AddCommand(CmdTransferUser())

	cmd.AddCommand(CmdCreateWhois())
	cmd.AddCommand(CmdUpdateWhois())
	cmd.AddCommand(CmdDeleteWhois())

	return cmd
}
