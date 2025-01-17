// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package model

// List represents a user-created list of accounts that the user follows.
//
// swagger:model list
type List struct {
	// The ID of the list.
	ID string `json:"id"`
	// The user-defined title of the list.
	Title string `json:"title"`
	// RepliesPolicy for this list.
	//	followed = Show replies to any followed user
	//	list = Show replies to members of the list
	//	none = Show replies to no one
	RepliesPolicy string `json:"replies_policy"`
}

// ListCreateRequest models list creation parameters.
//
// swagger:parameters listCreate
type ListCreateRequest struct {
	// Title of this list.
	// example: Cool People
	// in: formData
	// required: true
	Title string `form:"title" json:"title" xml:"title"`
	// RepliesPolicy for this list.
	//	followed = Show replies to any followed user
	//	list = Show replies to members of the list
	//	none = Show replies to no one
	// example: list
	// default: list
	// in: formData
	RepliesPolicy string `form:"replies_policy" json:"replies_policy" xml:"replies_policy"`
}

// ListUpdateRequest models list update parameters.
//
// swagger:parameters listUpdate
type ListUpdateRequest struct {
	// Title of this list.
	// example: Cool People
	// in: formData
	Title *string `form:"title" json:"title" xml:"title"`
	// RepliesPolicy for this list.
	//	followed = Show replies to any followed user
	//	list = Show replies to members of the list
	//	none = Show replies to no one
	// in: formData
	RepliesPolicy *string `form:"replies_policy" json:"replies_policy" xml:"replies_policy"`
}

// swagger:ignore
type ListAccountsChangeRequest struct {
	AccountIDs []string `form:"account_ids[]" json:"account_ids" xml:"account_ids"`
}
