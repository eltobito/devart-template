I have done some reseach to find which api is the best to retrieve data that are making sense for the thing I want to express. The Montreal Island image is now dead. I drop that idea.
I decide to compare 3 countries with data generate by people from those countries. Words that have a social meaning for me like love, fun or family.
I did'dnt decide yet the exact words and which service I gonna fetch.

To respect the rules of the project I start to look at Googles API. The one I found more interessting is  Youtube. Here are the kind of queries I tried.

https://www.googleapis.com/youtube/v3/search?part=snippet&publishedAfter=2013-12-21T00%3A00%3A00Z&publishedBefore=2014-03-20T00%3A00%3A00Z&q=democracy+ukraine&key={YOUR_API_KEY}

https://www.googleapis.com/youtube/v3/search?part=snippet&publishedAfter=2013-03-21T00%3A00%3A00Z&publishedBefore=2013-07-20T00%3A00%3A00Z&q=democratie+Quebec&key={YOUR_API_KEY}

But I'm not convince yet how these data are meaningfull

So for now I know I have to use an image library (PIL) and some call rest to retrieve data from social media (or kind of)  service.

