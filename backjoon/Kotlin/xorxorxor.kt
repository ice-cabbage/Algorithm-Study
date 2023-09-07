import java.io.BufferedReader
import java.io.BufferedWriter

private lateinit var bufferedReader: BufferedReader
private lateinit var bufferedWriter: BufferedWriter

fun main() {
    bufferedReader = System.`in`.bufferedReader()
    bufferedWriter = System.out.bufferedWriter()
	
    val (a, b, c) = bufferedReader
        .readLine()
        .split(" ")
        .map { it.toInt() }

    var answer = a
    
    if (c and 1 == 1) answer = answer xor b
    bufferedWriter.write("$answer")

    bufferedReader.close()
    bufferedWriter.close()
}
