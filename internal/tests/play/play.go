package play

import "encoding/xml"

type Play struct {
	XMLName xml.Name `xml:"PLAY"`
	Act     []struct {
		Epilogue struct {
			Speech struct {
				Line    []string `xml:"LINE"`
				Speaker string   `xml:"SPEAKER"`
			} `xml:"SPEECH"`
			StageDirection string `xml:"STAGEDIR"`
			Title          string `xml:"TITLE"`
		} `xml:"EPILOGUE"`
		Scene []struct {
			Speech []struct {
				Line []struct {
					CharData       string  `xml:",chardata"`
					StageDirection *string `xml:"STAGEDIR"`
				} `xml:"LINE"`
				Speaker        string    `xml:"SPEAKER"`
				StageDirection []*string `xml:"STAGEDIR"`
			} `xml:"SPEECH"`
			StageDirection []string `xml:"STAGEDIR"`
			Title          string   `xml:"TITLE"`
		} `xml:"SCENE"`
		Title string `xml:"TITLE"`
	} `xml:"ACT"`
	FrontMatter struct {
		Paragraph []string `xml:"P"`
	} `xml:"FM"`
	Personae struct {
		Persona      []string `xml:"PERSONA"`
		PersonaGroup []struct {
			GroupDescription string   `xml:"GRPDESCR"`
			Persona          []string `xml:"PERSONA"`
		} `xml:"PGROUP"`
		Title string `xml:"TITLE"`
	} `xml:"PERSONAE"`
	PlaySubtitle      string `xml:"PLAYSUBT"`
	ScreenDescription string `xml:"SCNDESCR"`
	Title             string `xml:"TITLE"`
}
