import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.LinkedList;
import java.util.Queue;

public class Main {

    static int[] bridge;
    static boolean[] visited;
    static int count;
    static int start;
    static int end;

    public static void main(String[] args) throws NumberFormatException, IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        count = Integer.parseInt(br.readLine());
        bridge = new int[count];
        visited = new boolean[count]; // 방문여부 확인

        String[] inputArray = br.readLine().split(" ");
        for(int i = 0; i < inputArray.length; i++){
            bridge[i] = Integer.parseInt(inputArray[i]);
        }

        String[] startEnd = br.readLine().split(" ");

        start = Integer.parseInt(startEnd[0])-1;
        end = Integer.parseInt(startEnd[1])-1;

        int minJump = bfs();

        if(minJump == Integer.MAX_VALUE) {
            bw.write(-1 + "\n");
        }
        else{
            bw.write(minJump + "\n");
        }

        bw.flush();
        bw.close();
        br.close();
    }

    private static int bfs() {

        Frog frog = new Frog(start, 0);
        Queue<Frog> needVisit = new LinkedList<>();
        needVisit.add(frog);

        int minJump = Integer.MAX_VALUE;

        while(!needVisit.isEmpty()){
            Frog now = needVisit.poll();
            visited[now.index] = true;

            if(now.index == end){
                if(now.jump < minJump){
                    minJump = now.jump;
                }

                break;
            }

            int multiple = 1;

            while(true){
                if(now.index + bridge[now.index] * multiple >= count){
                    break;
                }
                if(!visited[now.index + bridge[now.index] * multiple]){
                    Frog next = new Frog(now.index + bridge[now.index] * multiple, now.jump + 1);
                    needVisit.add(next);
                }
                multiple++;
            }

            multiple = 1;

            while(true){
                if(now.index - bridge[now.index] * multiple < 0){
                    break;
                }
                if(!visited[now.index - bridge[now.index] * multiple]){
                    Frog next = new Frog(now.index - bridge[now.index] * multiple, now.jump + 1);
                    needVisit.add(next);
                }
                multiple++;
            }
        }

        return minJump;
    }
}

class Frog {
    int index;
    int jump;

    public Frog(int index, int jump){
        this.index = index;
        this.jump = jump;
    }
}