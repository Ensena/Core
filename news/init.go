package news

import (
	"flag"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var client *twitter.Client

func init() {

	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "YLfn8Bs7uJvnKd9jWtvmTb6li", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "TpPeN62mgQJBu7k3hHuwzLZpoQmDlh3zHpZxsSvNQv91KbgSyP", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "76162357-U35OHSiIPLzVFS2S7iJdEzyJCTrOO3NWZ7xzM6Wj3", "Twitt Token")
	accessSecret := flags.String("access-secret", "zuXylnSBuz6owW3bwkJBBVKkmhhtaOWgkghPsLWtSE0vt", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client = twitter.NewClient(httpClient)

}
