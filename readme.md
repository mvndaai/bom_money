# Book of Mormon Money

This repo is to play with the money system found in [Alma 11](https://www.churchofjesuschrist.org/study/scriptures/bofm/alma/11?lang=eng) of the Book of Mormon

<br><hr><br>

Bases of money

* And the judge received for his wages according to his time—a senine of gold for a day
* a senine of gold [...] for a measure of barley, and also for a measure of every kind of grain.
* Zeezrom said unto [Amulek]: Behold, here are six onties of silver, and all these will I give thee if thou wilt deny the existence of a Supreme Being.

<br><hr><br>

A very interesting thing is that the money system seems to be binary based

| name | material | value | equation | binary |Book of Mormon description |
| - | - | - | - | - | - |
| Leah | Silver? | 1 | 2<sup>0</sup> | 000001 | a leah is the half of a shiblum |
| Shiblum | Silver? | 2 | 2<sup>1</sup> | 000010 | a shiblum is a half of a shiblon |
| Shiblon | Silver? | 4 | 2<sup>2</sup> | 000100 | a shiblon is half of a senum |
| Senum | Silver | 8 | 2<sup>3</sup> | 001000 | a senum of silver, which is equal to a senine of gold |
| Amnor | Silver | 16 | 2<sup>4</sup> | 010000 |an amnor of silver was as great as two senums |
| Ezrom | Silver | 32 | 2<sup>5</sup> | 100000 |an ezrom of silver was as great as four senums |
| Onti | Silver | 56 | 2<sup>3</sup>+2<sup>4</sup>+2<sup>5</sup> | 111000 | an onti was as great as them all |
| Antion | Gold | 12 | 2<sup>2</sup>+2<sup>3</sup> | 001100 | antion of gold is equal to three shiblons |
| Senine | Gold | 8 | 2<sup>3</sup> | 001000 | a senum of silver, which is equal to a senine of gold |
| Seon | Gold | 16 | 2<sup>4</sup> | 010000 | a seon of gold was twice the value of a senine |
| Shum | Gold | 32 | 2<sup>5</sup> | 100000 | a shum of gold was twice the value of a seon |
| Limnah | Gold | 56 | 2<sup>3</sup>+2<sup>4</sup>+2<sup>5</sup> | 111000 | a limnah of gold was the value of them all |

<br><hr><br>

Converting into US dollars gave me these values

```
** If a measure of barley cost the same as a 5lb Baking Flour: $13.00 **
$1.62 - silver leah
$3.25 - silver shiblum
$6.50 - silver shiblon
$13.00 - silver senum, gold senine
$19.50 - gold antion
$26.00 - silver amnor, gold seon
$52.00 - silver ezrom, gold shum
$91.00 - silver onti, gold limnah

** If a measure of barley cost the same as a congress person dalily wage: $478.02 **
$59.75 - silver leah
$119.50 - silver shiblum
$239.01 - silver shiblon
$478.02 - silver senum, gold senine
$717.03 - gold antion
$956.04 - silver amnor, gold seon
$1912.08 - silver ezrom, gold shum
$3346.14 - silver onti, gold limnah

** If a measure of barley cost the same as $1: $1.00 **
$0.12 - silver leah
$0.25 - silver shiblum
$0.50 - silver shiblon
$1.00 - silver senum, gold senine
$1.50 - gold antion
$2.00 - silver amnor, gold seon
$4.00 - silver ezrom, gold shum
$7.00 - silver onti, gold limnah

** If a measure of barley cost the same as $8 (leah the smallest coin = $1): $8.00 **
$1.00 - silver leah
$2.00 - silver shiblum
$4.00 - silver shiblon
$8.00 - silver senum, gold senine
$12.00 - gold antion
$16.00 - silver amnor, gold seon
$32.00 - silver ezrom, gold shum
$56.00 - silver onti, gold limnah
```