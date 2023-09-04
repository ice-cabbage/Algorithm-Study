import java.io.BufferedReader
import java.io.BufferedWriter
import java.io.InputStreamReader
import java.io.OutputStreamWriter
import java.util.*
data class xyz(val x:Int,val y:Int, val z:Int){
 
}
val br = BufferedReader(InputStreamReader(System.`in`))
val bw = BufferedWriter(OutputStreamWriter(System.out))
val line = readLine()!!.split(" ").map{it.toInt()}
val n = line[1]
val m = line[0]
val h = line[2]
 
val graph = mutableListOf<MutableList<MutableList<Int>>>()
val graph3 = mutableListOf<MutableList<Int>>()
val q : Queue<xyz> = LinkedList<xyz>()
val dx = mutableListOf<Int>(-1,1,0,0,0,0)
val dy = mutableListOf<Int>(0,0,-1,1,0,0)
val dz = mutableListOf<Int>(0,0,0,0,1,-1)
fun main()=with(br) {
 
    var graph2 = mutableListOf<MutableList<Int>>()
    for (i in 0 until h) {
        var graph2 = mutableListOf<MutableList<Int>>()
        for (j in 0 until n) {
            var go = readLine()!!.split(" ").map { it.toInt() }.toMutableList()
            graph2.add(go)
        }
        graph.add((graph2))
    }
 
    find()
    bfs()
 
    var key = true
    var sum = 0
    for(i in 0 until h){
        for(j in 0 until n) {
            for (k in 0 until m) {
                if (graph[i][j][k] == 0) {
                    key = false
                }
                if (sum < graph[i][j][k]) {
                    sum = graph[i][j][k]
                }
            }
        }
    }
    if(!key){
        bw.write("-1")
    }
    else{
        if(sum==1){
            bw.write("0")
        }
        else{
            bw.write("${sum-1}")
        }
    }
    bw.flush()
    bw.close()
}
 
fun find() {
    for(i in 0 until h){
        for(j in 0 until n) {
            for (k in 0 until m) {
                if (graph[i][j][k] == 1)
                    q.add(xyz(j,k,i))
            }
        }
    }
}
 
fun bfs(){
    while(!q.isEmpty()){
        var (a,b,c) = q.poll()
        for(i in 0 until 6){
            var dx = a+dx[i]
            var dy = b+dy[i]
            var dz = c+dz[i]
            if(dx<0 || dx>=n || dy<0 || dy>=m || dz<0 || dz>=h){
                continue
            }
            if(graph[dz][dx][dy]==-1){
                continue
            }
            if(graph[dz][dx][dy]==0){
                q.add(xyz(dx,dy,dz))
                graph[dz][dx][dy]=graph[c][a][b]+1
            }
        }
    }
}