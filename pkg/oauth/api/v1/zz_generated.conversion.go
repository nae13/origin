// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	api "github.com/openshift/origin/pkg/oauth/api"
	pkg_api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_ClusterRoleScopeRestriction_To_api_ClusterRoleScopeRestriction,
		Convert_api_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction,
		Convert_v1_OAuthAccessToken_To_api_OAuthAccessToken,
		Convert_api_OAuthAccessToken_To_v1_OAuthAccessToken,
		Convert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList,
		Convert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList,
		Convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken,
		Convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken,
		Convert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList,
		Convert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList,
		Convert_v1_OAuthClient_To_api_OAuthClient,
		Convert_api_OAuthClient_To_v1_OAuthClient,
		Convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization,
		Convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization,
		Convert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList,
		Convert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList,
		Convert_v1_OAuthClientList_To_api_OAuthClientList,
		Convert_api_OAuthClientList_To_v1_OAuthClientList,
		Convert_v1_ScopeRestriction_To_api_ScopeRestriction,
		Convert_api_ScopeRestriction_To_v1_ScopeRestriction,
	)
}

func autoConvert_v1_ClusterRoleScopeRestriction_To_api_ClusterRoleScopeRestriction(in *ClusterRoleScopeRestriction, out *api.ClusterRoleScopeRestriction, s conversion.Scope) error {
	out.RoleNames = in.RoleNames
	out.Namespaces = in.Namespaces
	out.AllowEscalation = in.AllowEscalation
	return nil
}

func Convert_v1_ClusterRoleScopeRestriction_To_api_ClusterRoleScopeRestriction(in *ClusterRoleScopeRestriction, out *api.ClusterRoleScopeRestriction, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleScopeRestriction_To_api_ClusterRoleScopeRestriction(in, out, s)
}

func autoConvert_api_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(in *api.ClusterRoleScopeRestriction, out *ClusterRoleScopeRestriction, s conversion.Scope) error {
	out.RoleNames = in.RoleNames
	out.Namespaces = in.Namespaces
	out.AllowEscalation = in.AllowEscalation
	return nil
}

func Convert_api_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(in *api.ClusterRoleScopeRestriction, out *ClusterRoleScopeRestriction, s conversion.Scope) error {
	return autoConvert_api_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(in, out, s)
}

func autoConvert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in *OAuthAccessToken, out *api.OAuthAccessToken, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = in.Scopes
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func Convert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in *OAuthAccessToken, out *api.OAuthAccessToken, s conversion.Scope) error {
	return autoConvert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in, out, s)
}

func autoConvert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in *api.OAuthAccessToken, out *OAuthAccessToken, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = in.Scopes
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func Convert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in *api.OAuthAccessToken, out *OAuthAccessToken, s conversion.Scope) error {
	return autoConvert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in, out, s)
}

func autoConvert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in *OAuthAccessTokenList, out *api.OAuthAccessTokenList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.OAuthAccessToken, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthAccessToken_To_api_OAuthAccessToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in *OAuthAccessTokenList, out *api.OAuthAccessTokenList, s conversion.Scope) error {
	return autoConvert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in, out, s)
}

func autoConvert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *api.OAuthAccessTokenList, out *OAuthAccessTokenList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthAccessToken, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthAccessToken_To_v1_OAuthAccessToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *api.OAuthAccessTokenList, out *OAuthAccessTokenList, s conversion.Scope) error {
	return autoConvert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in, out, s)
}

func autoConvert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in *OAuthAuthorizeToken, out *api.OAuthAuthorizeToken, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = in.Scopes
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func Convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in *OAuthAuthorizeToken, out *api.OAuthAuthorizeToken, s conversion.Scope) error {
	return autoConvert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in, out, s)
}

func autoConvert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *api.OAuthAuthorizeToken, out *OAuthAuthorizeToken, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	out.Scopes = in.Scopes
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func Convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *api.OAuthAuthorizeToken, out *OAuthAuthorizeToken, s conversion.Scope) error {
	return autoConvert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in, out, s)
}

func autoConvert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in *OAuthAuthorizeTokenList, out *api.OAuthAuthorizeTokenList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.OAuthAuthorizeToken, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in *OAuthAuthorizeTokenList, out *api.OAuthAuthorizeTokenList, s conversion.Scope) error {
	return autoConvert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in, out, s)
}

func autoConvert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *api.OAuthAuthorizeTokenList, out *OAuthAuthorizeTokenList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthAuthorizeToken, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *api.OAuthAuthorizeTokenList, out *OAuthAuthorizeTokenList, s conversion.Scope) error {
	return autoConvert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in, out, s)
}

func autoConvert_v1_OAuthClient_To_api_OAuthClient(in *OAuthClient, out *api.OAuthClient, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Secret = in.Secret
	out.AdditionalSecrets = in.AdditionalSecrets
	out.RespondWithChallenges = in.RespondWithChallenges
	out.RedirectURIs = in.RedirectURIs
	out.GrantMethod = api.GrantHandlerType(in.GrantMethod)
	if in.ScopeRestrictions != nil {
		in, out := &in.ScopeRestrictions, &out.ScopeRestrictions
		*out = make([]api.ScopeRestriction, len(*in))
		for i := range *in {
			if err := Convert_v1_ScopeRestriction_To_api_ScopeRestriction(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ScopeRestrictions = nil
	}
	return nil
}

func Convert_v1_OAuthClient_To_api_OAuthClient(in *OAuthClient, out *api.OAuthClient, s conversion.Scope) error {
	return autoConvert_v1_OAuthClient_To_api_OAuthClient(in, out, s)
}

func autoConvert_api_OAuthClient_To_v1_OAuthClient(in *api.OAuthClient, out *OAuthClient, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Secret = in.Secret
	out.AdditionalSecrets = in.AdditionalSecrets
	out.RespondWithChallenges = in.RespondWithChallenges
	out.RedirectURIs = in.RedirectURIs
	out.GrantMethod = GrantHandlerType(in.GrantMethod)
	if in.ScopeRestrictions != nil {
		in, out := &in.ScopeRestrictions, &out.ScopeRestrictions
		*out = make([]ScopeRestriction, len(*in))
		for i := range *in {
			if err := Convert_api_ScopeRestriction_To_v1_ScopeRestriction(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ScopeRestrictions = nil
	}
	return nil
}

func Convert_api_OAuthClient_To_v1_OAuthClient(in *api.OAuthClient, out *OAuthClient, s conversion.Scope) error {
	return autoConvert_api_OAuthClient_To_v1_OAuthClient(in, out, s)
}

func autoConvert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in *OAuthClientAuthorization, out *api.OAuthClientAuthorization, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.Scopes = in.Scopes
	return nil
}

func Convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in *OAuthClientAuthorization, out *api.OAuthClientAuthorization, s conversion.Scope) error {
	return autoConvert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in, out, s)
}

func autoConvert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *api.OAuthClientAuthorization, out *OAuthClientAuthorization, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.Scopes = in.Scopes
	return nil
}

func Convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *api.OAuthClientAuthorization, out *OAuthClientAuthorization, s conversion.Scope) error {
	return autoConvert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in, out, s)
}

func autoConvert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in *OAuthClientAuthorizationList, out *api.OAuthClientAuthorizationList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.OAuthClientAuthorization, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in *OAuthClientAuthorizationList, out *api.OAuthClientAuthorizationList, s conversion.Scope) error {
	return autoConvert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in, out, s)
}

func autoConvert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *api.OAuthClientAuthorizationList, out *OAuthClientAuthorizationList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthClientAuthorization, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *api.OAuthClientAuthorizationList, out *OAuthClientAuthorizationList, s conversion.Scope) error {
	return autoConvert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in, out, s)
}

func autoConvert_v1_OAuthClientList_To_api_OAuthClientList(in *OAuthClientList, out *api.OAuthClientList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.OAuthClient, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthClient_To_api_OAuthClient(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthClientList_To_api_OAuthClientList(in *OAuthClientList, out *api.OAuthClientList, s conversion.Scope) error {
	return autoConvert_v1_OAuthClientList_To_api_OAuthClientList(in, out, s)
}

func autoConvert_api_OAuthClientList_To_v1_OAuthClientList(in *api.OAuthClientList, out *OAuthClientList, s conversion.Scope) error {
	if err := pkg_api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := pkg_api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthClient, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthClient_To_v1_OAuthClient(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthClientList_To_v1_OAuthClientList(in *api.OAuthClientList, out *OAuthClientList, s conversion.Scope) error {
	return autoConvert_api_OAuthClientList_To_v1_OAuthClientList(in, out, s)
}

func autoConvert_v1_ScopeRestriction_To_api_ScopeRestriction(in *ScopeRestriction, out *api.ScopeRestriction, s conversion.Scope) error {
	out.ExactValues = in.ExactValues
	if in.ClusterRole != nil {
		in, out := &in.ClusterRole, &out.ClusterRole
		*out = new(api.ClusterRoleScopeRestriction)
		if err := Convert_v1_ClusterRoleScopeRestriction_To_api_ClusterRoleScopeRestriction(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ClusterRole = nil
	}
	return nil
}

func Convert_v1_ScopeRestriction_To_api_ScopeRestriction(in *ScopeRestriction, out *api.ScopeRestriction, s conversion.Scope) error {
	return autoConvert_v1_ScopeRestriction_To_api_ScopeRestriction(in, out, s)
}

func autoConvert_api_ScopeRestriction_To_v1_ScopeRestriction(in *api.ScopeRestriction, out *ScopeRestriction, s conversion.Scope) error {
	out.ExactValues = in.ExactValues
	if in.ClusterRole != nil {
		in, out := &in.ClusterRole, &out.ClusterRole
		*out = new(ClusterRoleScopeRestriction)
		if err := Convert_api_ClusterRoleScopeRestriction_To_v1_ClusterRoleScopeRestriction(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ClusterRole = nil
	}
	return nil
}

func Convert_api_ScopeRestriction_To_v1_ScopeRestriction(in *api.ScopeRestriction, out *ScopeRestriction, s conversion.Scope) error {
	return autoConvert_api_ScopeRestriction_To_v1_ScopeRestriction(in, out, s)
}
