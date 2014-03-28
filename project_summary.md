# Firearms or how to join the useless to the unpleasant


## Author
Name : Tobie Charette

Surname : Eltobito

GitHub : https://github.com/eltobito/devart-template


## Description
One of the most twisted relation we have in this modern society is the relation people have with firearms.
Most people live in cities. Why people have the need to own firearms? Why goverments are reluctant to 
limit the possesion of firearms like semi-automatic? This project want to give illustrate of the link betwen 
the abundance of firearm and the damage it created.

The final result is on my Picassa
<table style="width:194px;"><tr><td align="center" style="height:194px;background:url(https://www.gstatic.com/pwa/s/v/lighthousefe_213.01/transparent_album_background.gif) no-repeat left"><a href="https://picasaweb.google.com/114358707656696621914/DevArtOeuvreGunsControl?authuser=0&authkey=Gv1sRgCIfAvryl_6XPqgE&feat=embedwebsite"><img src="https://lh4.googleusercontent.com/-nXdgHWNnSsw/UzTpKakr8AE/AAAAAAAAMIM/V1PdzIawKJM/s160-c/DevArtOeuvreGunsControl.jpg" width="160" height="160" style="margin:1px 0 0 4px;"></a></td></tr><tr><td style="text-align:center;font-family:arial,sans-serif;font-size:11px"><a href="https://picasaweb.google.com/114358707656696621914/DevArtOeuvreGunsControl?authuser=0&authkey=Gv1sRgCIfAvryl_6XPqgE&feat=embedwebsite" style="color:#4D4D4D;font-weight:bold;text-decoration:none;">DevArt - Oeuvre Guns Control</a></td></tr></table>

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
[gunpolicy.org](http://www.gunpolicy.org/fr/documents "Gun reference")
and
[Wikipedia](http://en.wikipedia.org/wiki/List_of_countries_by_firearm-related_death_rate "Gun reference")






