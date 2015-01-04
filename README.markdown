
## About

"You're late.  And this is the most important test of your life.  Does anyone
know what 2 + 2 equals?"

"5?"

"Correct.  Anyone else?"

"6.  6 minutes left."

"Correct.  Simon?"

"Uh, 4."

"And exactly what formula did _you_ use?"

## Thoughts

This is going to be a "currency exchange" simulation with a very faulty
database.  I don't know how a real currency exchange works though, so it won't
be a particularly accurate simulation.  I'm really not sure how this is going to
turn out, but here's what I'm thinking now:

- On each turn, players decide what trades to make, and they are all executed at
once (at the end of the turn).

- Players will control hedge funds.  There will be a few "banks" that you can
always make fair trades with.  There will be exchange rates between each
currency - these rates might fluctuate from turn to turn.

- Hopefully, the fun part will be to try construct trades that cause favorable
data inconsistency.

- I'm really unsure about the objective of the game.  I don't think it's
zero-sum, since there's technically no resource that is limited.  So the
players could try to team up to "generate" the most money in the system.  Or
they could work together to beat the banks (whatever that means).  Or they
could just try to maximize their own hedge fund's accounts.
