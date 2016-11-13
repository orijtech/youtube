package youtube_test

import (
	"fmt"
	"log"

	"github.com/odeke-em/youtube"
)

func ExampleSearch() {
	client, err := youtube.New()
	if err != nil {
		log.Fatal(err)
	}

	param := &youtube.SearchParam{
		Query:             "James Bond Best scenes",
		MaxPage:           2,
		MaxResultsPerPage: 2,
	}

	pagesChan, err := client.Search(param)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Query: %q\n", param.Query)

	for page := range pagesChan {
		if page.Err != nil {
			log.Fatal(page.Err)
		}
		for _, item := range page.Items {
			snippet := item.Snippet
			videoId := item.Id.VideoId
			fmt.Printf("Title: %q Description: %q Id: %q\n", snippet.Title, snippet.Description, videoId)
		}
	}

	// Output:
	// Query: "James Bond Best scenes"
	// Title: "Top 10 James Bond Moments" Description: "Shaken or stirred, Craig or Brosnan, ejector seat or exploding pen - whichever way you look at it, there is no denying that 007 is one of the coolest and most ..." Id: "Bpcupsxy0T8"
	// Title: "Skyfall - Glass Marksman Shot (1080p)" Description: "The scene in the James Bond movie \"Skyfall\" where Bond and Silva shoots at a glass placed on Severines head." Id: "5Z0wEVxL30Q"
	// Title: "BEST SCENES OF JAMES BOND SPECTRE 007" Description: "" Id: "4BpFHOB1q04"
	// Title: "The Best James Bond Car Chase Scenes" Description: "A compilation of the Best James Bond Car Chase Scenes! Try guessing the films if you are a James Bond fan ;) Subscribe to The Wheels and Chips Journal!" Id: "uS9BFHgNQag"
}

func ExampleSearchRelatedToVideo() {
	client, err := youtube.New()
	if err != nil {
		log.Fatal(err)
	}

	param := &youtube.SearchParam{
		MaxPage:           1,
		MaxResultsPerPage: 10,
		RelatedToVideo:    "vuUIoPVRpmk",
	}

	pagesChan, err := client.Search(param)
	if err != nil {
		log.Fatal(err)
	}

	for page := range pagesChan {
		if page.Err != nil {
			log.Fatal(page.Err)
		}
		for _, item := range page.Items {
			snippet := item.Snippet
			videoId := item.Id.VideoId
			fmt.Printf("Title: %q Description: %q Id: %q\n", snippet.Title, snippet.Description, videoId)
		}
	}
}

func ExampleMostPopular() {
	client, err := youtube.New()
	if err != nil {
		log.Fatal(err)
	}

	param := &youtube.SearchParam{
		MaxPage:           1,
		MaxResultsPerPage: 10,
	}

	videoPages, err := client.MostPopular(param)
	if err != nil {
		log.Fatal(err)
	}

	for videoPage := range videoPages {
		if videoPage.Err != nil {
			log.Fatal(videoPage.Err)
		}

		for _, video := range videoPage.Items {
			snippet := video.Snippet
			stats := video.Statistics
			fmt.Printf("\nViewCount: %d Title: %q Description: %q Id: %q\n",
				stats.ViewCount, snippet.Title, snippet.Description, video.Id)
		}
	}
}

func ExampleById() {
	client, err := youtube.New()
	if err != nil {
		log.Fatal(err)
	}

	videoPages, err := client.ById("Bpcupsxy0T8", "FLeMssD0R3Y", "Jv95aptVSUk", "rTAZlHGOVo8")
	if err != nil {
		log.Fatal(err)
	}

	for videoPage := range videoPages {
		if videoPage.Err != nil {
			log.Fatal(videoPage.Err)
		}

		for _, video := range videoPage.Items {
			snippet := video.Snippet
			fmt.Printf("\nTitle: %q Description: %q Id: %q\n", snippet.Title, snippet.Description, video.Id)
		}
	}

	// Output:
	//
	// Title: "Top 10 James Bond Moments" Description: "Shaken or stirred, Craig or Brosnan, ejector seat or exploding pen - whichever way you look at it, there is no denying that 007 is one of the coolest and most iconic secret agents of all time. Join http://www.WatchMojo.com as we count down our picks for the Top 10 James Bond Moments. For this list, we are focusing on all films and iconic sequences from the famous franchise involving our British secret serviceman. Click here to subscribe: http://www.youtube.com/subscription_center?add_user=watchmojo or visit our channel page here: http://www.youtube.com/watchmojo Also, check out our interactive Suggestion Tool at http://www.WatchMojo.com/suggest :)\n\nCheck us out at http://www.Twitter.com/WatchMojo, http://instagram.com/watchmojo and http://www.Facebook.com/WatchMojo. \n\nSpecial thanks to our users WatchDogsFan47, derfboy00, Anthony Redford, Andrew A. Dennison, Eduardo Silva and Tommy Carr for submitting the idea on our Interactive Suggestion Tool at http://www.WatchMojo.com/suggest\n\nCheck out the voting page here, \nhttp://www.watchmojo.com/suggest/Top+10+Best+James+Bond+Moments\n\nWant a WatchMojo cup, mug, t-shirts, pen, sticker and even a water bottle?  Get them all when you order your MojoBox gift set here:\nhttp://watchmojo.com/store/\n\nWatchMojo is a leading producer of reference online video content, covering the People, Places and Trends you care about.\nWe update DAILY with 4-5 Top 10 lists, Origins, Biographies, Versus clips on movies, video games, music, pop culture and more!" Id: "Bpcupsxy0T8"
	//
	// Title: "Desiigner - Zombie Walk ft. King Savage" Description: "“Zombie Walk\" Official Music Video From Desiigner’s Debut Mixtape, New English\nDirected by Grant Curatola\nProduced by Taylor Shung \nCinematography by Eric K. Yue. \nEdited by Zoe Mougin \nAnimation by Pixel Pirate Studio \nLine Produced by Lizzie Shapiro\nColor by Mikey Rossiter \nSound Design by Jeff Malen\n\nStream/Download \"New English\"\nDownload: http://smarturl.it/dNewEnglish \nStream: http://smarturl.it/sNewEnglish \n\nMore From Desiigner: \nhttp://lifeofdesiigner.com \nhttp://twitter.com/lifeofdesiigner\nhttp://facebook.com/lifeofdesiigner\nhttp://instagram.com/lifeofdesiigner\nhttp://youtube.com/DesiignerVEVO\n\nMusic video by Desiigner performing Zombie Walk. (C) 2016 Getting Out Our Dreams, Inc./Def Jam Recordings, a division of UMG Recordings, Inc.\n\nhttp://vevo.ly/fG3AvT" Id: "FLeMssD0R3Y"
	//
	// Title: "Anderson .Paak - The Dreamer (feat. Talib Kweli & Timan Family Choir)" Description: "From the album \"Malibu\". Out Now!\niTunes: https://itunes.apple.com/us/album/malibu/id1065681363&app=itunes\nSpotify: http://open.spotify.com/album/4VFG1DOuTeDMBjBLZT7hCK\nAmazon: http://www.amazon.com/s/ref=nb_sb_noss?url=search-alias%3Ddigital-music&field-keywords=Malibu\nGoogle Play: https://play.google.com/store/music/album/Anderson_Paak_Malibu?id=B562fkkr2nawzqrty3v5vr2x7fe\n\nOBE / Steel Wool / Art Club / EMPIRE" Id: "Jv95aptVSUk"
	//
	// Title: "Wiz Khalifa - Bake Sale ft. Travis Scott [Official Video]" Description: "See Wiz & Snoop on tour this summer http://smarturl.it/TheHighRoadTour\n\n2 million views and I'll drop another one. #2millionTaylors\n\nGet Taylor Gang merch http://store.taylorgang.com\n\nDownload \"Khalifa\" http://smarturl.it/KHALIFA\nStream \"Khalifa\" http://smarturl.it/KHALIFAspotify\nStream \"Khalifa\" http://smarturl.it/KHALIFAapplemusic\n\nDirected by Dan Folger x Matt Meehan x Wiz Khalifa\n\n►Subscribe to channel: http://goo.gl/y3Bnno\n►Snapchat - https://www.snapchat.com/add/khalifathecap\n►Twitter - https://twitter.com/wizkhalifa\n►Facebook - https://facebook.com/wizkhalifa\n►Instagram - https://instagram.com/wizKhalifa\n►Soundcloud - https://soundcloud.com/wizkhalifa\n►Website: http://wizkhalifa.com\n\nWiz Khalifa - Bake Sale ft. Travis Scott [Official Video]" Id: "rTAZlHGOVo8"
}
