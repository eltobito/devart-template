I have done some research to find which api is the best to retrieve data that are making sense for the feeling I want to expose. The Montreal Island image is now dead. I drop that idea.

I decide to compare 3 countries with data generate by people from those countries. Words that have a social meaning for me like love, fun or family.
I didn't decide yet the exact words and which service I'm gonna fetch.

To respect the rules of the project,I start to look at Google's APIs. The one I found more interessting is  Youtube. Here are the kind of queries I tried.

https://www.googleapis.com/youtube/v3/search?part=snippet&publishedAfter=2013-12-21T00%3A00%3A00Z&publishedBefore=2014-03-20T00%3A00%3A00Z&q=democracy+ukraine&key={YOUR_API_KEY}

https://www.googleapis.com/youtube/v3/search?part=snippet&publishedAfter=2013-03-21T00%3A00%3A00Z&publishedBefore=2013-07-20T00%3A00%3A00Z&q=democratie+Quebec&key={YOUR_API_KEY}

But I'm not convince yet how these data are meaningfull for my project.

So for now I know I have to use an image library (PIL) and some call rest to retrieve data from social media (or kind of)  service.

Example of JSON I'll have to parse

    - Show headers -
      
    {
     "kind": "youtube#searchListResponse",
     "etag": "\"PoAnP6GALHRQwbIPTryP-ZGqRQg/r7u4v_cJ2xETj8TM-QbTYe_b6ZA\"",
     "nextPageToken": "CAUQAA",
     "pageInfo": {
      "totalResults": 1000000,
      "resultsPerPage": 5
     },
     "items": [
      {
       "kind": "youtube#searchResult",
       "etag": "\"PoAnP6GALHRQwbIPTryP-ZGqRQg/HxtOyVQ7kbYW9oJUw2xT-pFYkFc\"",
       "id": {
        "kind": "youtube#video",
        "videoId": "Wp3zyoOROUo"
       },
       "snippet": {
        "publishedAt": "2014-03-17T14:42:25.000Z",
        "channelId": "UCPXRVaGxKOgHXXubivY4-nw",
        "title": "My Transition to Eating 100% RAW ♥ (Skin & Health Improvements!)",
        "description": "Hi everyone! This video was initially A LOT longer because I guess there's just so much to say          on this \"new change\" :) but I tried to keep all the topics sho...",
