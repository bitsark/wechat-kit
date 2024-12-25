package offiaccount

import "encoding/xml"

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html

const (
	MessageTypeText       = "text"
	MessageTypeImage      = "image"
	MessageTypeVoice      = "voice"
	MessageTypeVideo      = "video"
	MessageTypeShortVideo = "shortvideo"
	MessageTypeLocation   = "location"
	MessageTypeLink       = "link"
)

type (
	ToUserName struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"ToUserName"`
	}

	FromUserName struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"FromUserName"`
	}

	MessageType struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"MsgType"`
	}

	MessageContent struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Content"`
	}

	PicURL struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"PicUrl"`
	}

	MediaID struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"MediaId"`
	}

	Recognition struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Recognition"`
	}

	Format struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Format"`
	}

	ThumbMediaID struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"ThumbMediaId"`
	}

	Label struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Label"`
	}

	Title struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Title"`
	}

	Description struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Description"`
	}

	URL struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"Url"`
	}

	MusicURL struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"MusicUrl"`
	}

	HQMusicURL struct {
		Text    string   `xml:",cdata"`
		XMLName xml.Name `xml:"HQMusicUrl"`
	}
)

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html
type (
	CiphertMessageTypeReq struct {
	}

	ReqMessageBase struct {
		XMLName      xml.Name     `xml:"xml"`
		ToUserName   ToUserName   `xml:"ToUserName"`
		FromUserName FromUserName `xml:"FromUserName"`
		CreateTime   int64        `xml:"CreateTime"`
		MessageID    uint64       `xml:"MsgId"`
		MessageType  MessageType  `xml:"MsgType"`
	}

	TextReqMessage struct {
		ReqMessageBase
		MessageContent MessageContent `xml:"Content"`
	}

	ImageReqMessage struct {
		ReqMessageBase
		PicURL  PicURL  `xml:"PicUrl"`
		MediaID MediaID `xml:"MediaId"`
	}

	VoiceReqMessage struct {
		ReqMessageBase
		Format      Format      `xml:"Format"`
		MediaID     MediaID     `xml:"MediaId"`
		Recognition Recognition `xml:"Recognition"`
	}

	VideoReqMessage struct {
		ReqMessageBase
		Format       Format       `xml:"Format"`
		MediaID      MediaID      `xml:"MediaId"`
		ThumbMediaID ThumbMediaID `xml:"ThumbMediaId"`
	}

	ShortVideoReqMessage VoiceReqMessage

	LocationReqMessage struct {
		ReqMessageBase
		Label     Label   `xml:"Label"`
		Scale     int     `xml:"Scale"`
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
	}

	Link struct {
		ReqMessageBase
		URL         URL         `xml:"Url"`
		Title       Title       `xml:"Title"`
		Description Description `xml:"Description"`
	}
)

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html
type (
	RespMessageBase struct {
		XMLName      xml.Name     `xml:"xml"`
		ToUserName   ToUserName   `xml:"ToUserName"`
		FromUserName FromUserName `xml:"FromUserName"`
		CreateTime   int64        `xml:"CreateTime"`
		MessageType  MessageType  `xml:"MsgType"`
	}

	TextRespMessage struct {
		RespMessageBase
		MessageContent MessageContent `xml:"Content"`
	}

	Image struct {
		MediaID MediaID `xml:"MediaId"`
	}
	ImageRespMessage struct {
		ReqMessageBase
		Image Image `xml:"Image"`
	}

	Voice struct {
		MediaID MediaID `xml:"MediaId"`
	}
	VoiceRespMessage struct {
		RespMessageBase
		Voice Voice `xml:"Voice"`
	}

	Video struct {
		Title       Title       `xml:"Title"`
		MediaID     MediaID     `xml:"MediaId"`
		Description Description `xml:"Description"`
	}
	VideoRespMessage struct {
		ReqMessageBase
		Video Video `xml:"Video"`
	}

	Music struct {
		Title        Title        `xml:"Title"`
		Description  Description  `xml:"Description"`
		MusicURL     MusicURL     `xml:"MusicUrL"`
		HQMusicURL   HQMusicURL   `xml:"HQMusicUrL"`
		ThumbMediaID ThumbMediaID `xml:"ThumbMediaId"`
	}
	MusicRespMessage struct {
		RespMessageBase
		Music Music `xml:"Music"`
	}

	Article struct {
		Title       Title       `xml:"Title"`
		Description Description `xml:"Description"`
		PicURL      PicURL      `xml:"PicUrl"`
		URL         URL         `xml:"Url"`
	}
	ArticleRespMessage struct {
		RespMessageBase
		ArticleCount int       `xml:"ArticleCount"`
		Articles     []Article `xml:"Articles>item"`
	}
)
