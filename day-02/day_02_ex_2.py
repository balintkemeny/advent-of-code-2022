matchupValues = {
    "X": 0,
    "Y": 3,
    "Z": 6,
}

shapeGuide = {
    "A X": 3,
    "A Y": 1,
    "A Z": 2,
    "B X": 1,
    "B Y": 2,
    "B Z": 3,
    "C X": 2,
    "C Y": 3,
    "C Z": 1,
}

totalScore = 0

with open("input.txt", "r") as f:
    for line in f.readlines():
        totalScore += shapeGuide[line.rstrip()]
        totalScore += matchupValues[line[2]]        

print(totalScore)