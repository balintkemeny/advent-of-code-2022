shapeValues = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}

matchups = {
    "A X": 3,
    "A Y": 6,
    "A Z": 0,
    "B X": 0,
    "B Y": 3,
    "B Z": 6,
    "C X": 6,
    "C Y": 0,
    "C Z": 3,
}

totalScore = 0

with open("input.txt", "r") as f:
    for line in f.readlines():
        totalScore += shapeValues[line[2]]
        totalScore += matchups[line.rstrip()]

print(totalScore)