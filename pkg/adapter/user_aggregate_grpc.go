// Copyright 2020 The Alexandria Foundation
//
// Licensed under the GNU Affero General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"github.com/alexandria-oss/identity-api/internal/domain/aggregate"
	"github.com/alexandria-oss/identity-api/proto"
)

// UserToGRPC Convert a user aggregate root to gRPC struct
func UserToGRPC(ag *aggregate.UserRoot) *proto.User {
	// Avoid nil refs
	prefUsername := ""
	if ag.User.PreferredUsername != nil {
		prefUsername = *ag.User.PreferredUsername
	}

	midName := ""
	if ag.User.MiddleName != nil {
		midName = *ag.User.MiddleName
	}

	famName := ""
	if ag.User.FamilyName != nil {
		famName = *ag.User.FamilyName
	}

	picture := ""
	if ag.User.Picture != nil {
		picture = *ag.User.Picture
	}

	return &proto.User{
		Id:                ag.User.ID.Value,
		Username:          ag.User.Username.Value,
		PreferredUsername: prefUsername,
		Email:             ag.User.Email.Value,
		Name:              ag.User.Name,
		MiddleName:        midName,
		FamilyName:        famName,
		Locale:            ag.User.Locale,
		Picture:           picture,
		Status:            ag.User.Status,
		CreateTime:        ag.User.CreateTime.String(),
		UpdateTime:        ag.User.UpdateTime.String(),
		Enabled:           ag.User.Enabled.Value,
	}
}

// BulkUserToGRPC Convert a user aggregate root's slice to gRPC-only slice
func BulkUserToGRPC(agSlice []*aggregate.UserRoot) []*proto.User {
	users := make([]*proto.User, 0)
	for _, u := range agSlice {
		users = append(users, UserToGRPC(u))
	}

	return users
}
