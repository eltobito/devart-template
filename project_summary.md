# Firearms


## Author
Name : Tobie Charette

Surname : Eltobito

GitHub : https://github.com/eltobito/devart-template


## Description
One of the most twisted relation we have in this modern society is the relation people have with firearms.
Most people live in cities. Why people have the need to own firearms? Why goverments are reluctant to 
limit the possesion of firearms like semi-automatic? This project want give an overview of the presence of
firearms in our society



## Example Code

```
func prepareData(pays Country, pop int) (int, int, int) {

	var deathbygun = pays.Deathbygun * float64(pop)
	var guns = pays.Guns * float64(pop)
	var rank = int(pays.Rank * 255 / 178)

	return int(deathbygun), int(guns), rank
}
```
## Links to External Libraries

[Draw2d](https://code.google.com/p/draw2d/ "Draw2d")

[TOML](https://github.com/BurntSushi/toml "TOML")



##Data reference
[Gun reference](http://www.gunpolicy.org/fr/documents "Gun reference")

## Images & Videos
NOTE: For additional images you can either use a relative link to an image on this repo or an absolute link to an externally hosted image.




