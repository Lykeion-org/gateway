package handlers

import (
	"github.com/Lykeion-org/gateway/internal/models"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
)

func ProtoToReferent(r *shared.Referent) *models.Referent {
	if r == nil {
		return nil
	}

	referent := &models.Referent{
		Uid:         r.GetUid(),
		EnReference: r.GetEnReference(),
		ImageSource: derefString(stringPtr(r.GetImageSource())),
		Symbols:     ProtoToSymbolList(r.GetSymbols()),
	}
	return referent
}

func ProtoToReferentList(r []*shared.Referent) []models.Referent {
	if r == nil {
		return []models.Referent{}
	}

	result := make([]models.Referent, len(r))
	for i, referent := range r {
		if referent != nil {
			result[i] = *ProtoToReferent(referent)
		}
	}
	return result
}

func ProtoToSymbol(r *shared.Symbol) *models.Symbol {
	if r == nil {
		return nil
	}

	symbol := &models.Symbol{
		Uid:        r.GetUid(),
		Language:   int(r.GetLanguage()),
		SymbolType: int(r.GetSymbolType()),
	}

	if r.Lemma != nil {
		symbol.Lemma = *ProtoToWord(r.GetLemma())
	}

	if r.Words != nil {
		symbol.Words = ProtoToWordList(r.GetWords())
	} else {
		symbol.Words = []models.Word{}
	}

	return symbol
}

func ProtoToSymbolList(r []*shared.Symbol) []models.Symbol {
	if r == nil {
		return []models.Symbol{}
	}

	symbols := make([]models.Symbol, len(r))
	for i, symbol := range r {
		if symbol != nil {
			symbols[i] = *ProtoToSymbol(symbol)
		}
	}
	return symbols
}

func ProtoToWord(req *shared.Word) *models.Word {
	if req == nil {
		return nil
	}

	word := &models.Word{
		Uid:         req.GetUid(),
		Word:        req.GetWord(),
		Ipa:         req.GetIpa(),
		SoundSource: req.GetSoundSource(),
	}

	if req.Attributes != nil {
		word.Attributes = ProtoToWordAttributes(req.GetAttributes())
	}

	return word
}

func ProtoToWordAttributes(a *shared.WordAttributes) models.WordAttributes {
	return models.WordAttributes{
		Gender:         int(a.GetGender()),
		Number:         int(a.GetNumber()),
		Unique:         a.GetUnique(),
		Diminutive:     a.GetDiminutive(),
		Case:           int(a.GetCase()),
		Mood:           int(a.GetMood()),
		Tense:          int(a.GetTense()),
		Aspect:         int(a.GetAspect()),
		Person:         int(a.GetPerson()),
		DirectObject:   a.GetDirectObject(),
		IndirectObject: a.GetIndirectObject(),
		Valency:        int(a.GetValency()),
		Reflexive:      a.GetReflexive(),
	}
}

func ProtoToWordSuggestion(s *lexora.WordSuggestion) models.WordSuggestion {
	return models.WordSuggestion{
		Uid:          s.GetUid(),
		Word:         s.GetWord(),
		Lemma:        s.GetLemma(),
		Language:     int(s.GetLanguage()),
		SymbolType:   int(s.GetSymbolType()),
		WordType:     int(s.GetWordType()),
		Attributes:   ProtoToWordAttributes(s.GetAttributes()),
		JobId:        s.GetJobUid(),
		RequesterUid: s.GetRequesterUid(),
	}
}

func ProtoToWordList(words []*shared.Word) []models.Word {
	if words == nil {
		return []models.Word{}
	}

	result := make([]models.Word, len(words))
	for i, word := range words {
		if word != nil {
			result[i] = *ProtoToWord(word)
		}
	}
	return result
}

func derefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}