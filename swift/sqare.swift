import Foundation

func countSquares() {

    guard let nStr = readLine(), let n = Int(nStr) else {
        print("Number count input error")
        return
    }

    guard let numbersStr = readLine() else {
        print("Number imput error")
        return
    }

    let numbers = numbersStr.split(separator: " ").compactMap { Int($0) }

    guard numbers.count == n else {
        print("Number count mismatch")
        return
    }

    var count = 0

    for number in numbers {
        let root = Int(sqrt(Double(number)))
        if root * root == number {
            count += 1
        }
    }
    print(count)
}

countSquares()
