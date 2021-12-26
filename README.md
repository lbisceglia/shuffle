# 99 Card Game &#127145;&#127161;

## Gameplay
99 is a turn-based, multiplayer card game. Players are each dealt three cards. As the play goes around, each player puts down one card, adding its total to the running total. If the running total exceeds 99, the round is over and the player who forced the total over 99 is given a strike. Three strikes, you're out! Last player standing wins.

The point total starts at zero and each card adds its face value in points (e.g. a 5 is worth five points, a face card is worth 10 points) except for certain cards that have special values or meanings:
* 4 reverses play (and is worth 0 points)
* 10 is worth -10 points
* K is worth 0 points
* 9 takes the point total straight to 99 (or keeps it 99 if the total is already 99)

Watch out for those 9's!

Each player draws a replacement card after they play.

## Getting Started
1. [Install Go](https://golang.org/doc/install)
2. Clone the repo
3. In a terminal, run `go run .` to play the command line game

Future versions will be containerized to avoid installing Go or other dependencies locally.

## Implementation Notes
The current implementation is a command line proof of concept that runs in the terminal. It plays a single 4-person round of the game, with all cards visible for illustrative purposes.

Future improvements will enable true multiplayer games with multiple clients. These may include:
* Refactoring the game into a simple client-server web app or mobile app
* Leveraging the concurrency features of Go to handle concurrent player requests on the deck (e.g. drawing cards at the same time)
* Using WebSockets (or similar) to keep players updated on the running count

## Inspiration
This game is a family favourite that has become a Christmas day tradition. After a hearty Christmas dinner at Nonna's house, we push the tables end-to-end, cobble together every last deck of cards we can scrounge, throw in five dollars apiece, and gather round for an epic forty-person 99 showdown. With 120 rounds ahead of us, we buckle up for a thrilling game that lasts all night and crowns one champion, pausing only for reloads of wine, taralli, salami, and cheese. There only one question left...

Who's gonna shuffle?
