import Data.List

matrixElementsSum :: [[Int]] -> Int
matrixElementsSum xss = sum [x | xs <- (transpose xss), x <- takeWhile (> 0) xs]

testData :: [(String, [[Int]], Int)]
testData = [ ( "one"
             , [[0, 1, 1, 2], [0, 5, 0, 0], [2, 0, 3, 3]]
             ,  9 )
           , ( "two"
             , [[1, 1, 1, 0], [0, 5, 0, 1], [2, 1, 3, 10]]
             , 9 ) ]

main = print (all (== True) [matrixElementsSum matrix == expected | (_, matrix, expected) <- testData])
