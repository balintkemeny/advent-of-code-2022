shapeValues = {
    "X": 1,
    "Y": 2,
    "Z": 3,
    "A": 1,
    "B": 2,
    "C": 3,
}
matchResultPointValues = [0, 3, 6]

totalScore = 0

with open("input.txt", "r") as f:
    for line in f.readlines():
        ourShapeValue = shapeValues[line[2]]
        opponentShapeValue = shapeValues[line[0]]
        totalScore += ourShapeValue + matchResultPointValues[(ourShapeValue - opponentShapeValue + 1) % 3]

print(totalScore)