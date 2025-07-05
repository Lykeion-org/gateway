package models

type User struct {
	Uid					string			`json:"uid"`			
	EmailAddress		string			`json:"emailAddress"`		
	AvatarSource		string			`json:"avatarSource"`		
	NativeLanguage		int32			`json:"nativeLanguage"`	
	ReferenceLanguage	int32			`json:"referenceLanguage"`	
	TargetLanguage		int32			`json:"targetLanguage"`	
	FirstName			string			`json:"firstName"`		
	LastName			string			`json:"lastName"`		
	ExtraName			string			`json:"extraName"`		
	DisplayName			string			`json:"displayName"`		
	Role				int32			`json:"role"`	
}