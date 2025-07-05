package models


type Referent struct {
	Uid 			string 			`json:"uid"`
	EnReference 	string 			`json:"enReference"`
	ImageSource 	string 			`json:"imageSource"`
	Symbols 		[]Symbol 		`json:"symbols"`
}

type Symbol struct {
	Uid 			string 			`json:"uid"`
	Language 		int 			`json:"language"`
	SymbolType 		int 			`json:"symbolType"`
	Lemma 			Word 			`json:"lemma"`
	Words 			[]Word 			`json:"words"`
}

type Word struct {
	Uid 			string 			`json:"uid"`
	Word 			string 			`json:"word"`
	Ipa 			string 			`json:"ipa"`
	SoundSource 	string 			`json:"soundSource"`
	Attributes 		WordAttributes 	`json:"attributes"`
}

type WordAttributes struct {
	Gender   		int 			`json:"gender"`
	Number 			int				`json:"number"`
	Unique 			bool			`json:"uniqe"`
	Diminutive 		bool			`json:"diminutive"`
	Case 			int				`json:"case"`
	Mood 			int				`json:"mood"`
	Tense 			int				`json:"tense"`
	Aspect 			int				`json:"aspect"`
	Person 			int				`json:"person"`
	DirectObject 	bool			`json:"directObject"`
	IndirectObject 	bool			`json:"indirectObject"`
	Valency 		int				`json:"valency"`
	Reflexive 		bool			`json:"reflexive"`
}

type CreateReferentRequest struct {
	EnReference 	string 			`json:"enReference" binding:"required"`
	ImageSource 	string 			`json:"imageSource"`
}

type CreateSymbolRequest struct {
	ReferentUid 	string 			`json:"referentUid"`
    Language 		int 			`json:"language"`
    SymbolType 		int 			`json:"symbolType"`
}

type CreateWordRequest struct {
	SymbolUid 		string 			`json:"symbolUid"`
	Word 			string 			`json:"word" binding:"required"`
	SoundSource 	string 			`json:"soundSource"`
	Ipa 			string 			`json:"ipa"`
	WordType 		int 			`json:"wordType"`
}

type LinkSymbolToReferentRequest struct {
	SymbolUid 		[]string 		`json:"symbolUid" binding:"required"`
	ReferentUid 	string 			`json:"referentUid" binding:"required"`
}

type LinkWordToSymbolRequest struct {
	SymbolUid 		string 			`json:"symbolUid" binding:"required"`
	WordUid 		[]string 		`json:"wordUid" binding:"required"`
}

type SetWordAsLemmaRequest struct {
	SymbolUid 		string 			`json:"symbolUid" binding:"required"`
	WordUid 		string 			`json:"wordUid" binding:"required"`
}

type AnalyzeTextRequestInternal struct {
	JobId 			string 			`json:"job_id"`
	RequesterId 	string 			`json:"requester_id"`
	Text 			string 			`json:"text"`
}

type AnalyzeTextRequestClient struct {
	Text 			string 			`json:"text"`
}

type AnalyzeTextResponse struct {
	JobId 			string 			`json:"jobID"`
}