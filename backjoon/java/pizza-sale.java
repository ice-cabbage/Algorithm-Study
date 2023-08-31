import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int k = Integer.parseInt(br.readLine()); 
        String[] x =  br.readLine().split(" ");
        int m = Integer.parseInt(x[0]);
        int n = Integer.parseInt(x[1]);
        int[] arrA = new int[m];
        int[] arrB = new int[n];
        int sumA = 0;
        int sumB = 0;
        for (int i = 0; i < m; i++) {
            arrA[i] = Integer.parseInt(br.readLine());
            sumA +=arrA[i];
        }
        for (int i = 0; i < n; i++) {
            arrB[i] = Integer.parseInt(br.readLine());
            sumB +=arrB[i];
        }

        int[] DP1 = new int[k+1];
        int[] DP2 = new int[k+1];
        DP1[0] = 1;
        DP2[0] = 1;
        count_pieces(DP1,arrA,k);
        count_pieces(DP2,arrB,k);        
        sumA = sumA > k ? 0 : sumA;
        sumB = sumB > k ? 0 : sumB;
        DP1[sumA]=1;
        DP2[sumB]=1;

        int res=0;
        for (int i = 0; i <= k; i++)
            res += DP1[i] * DP2[k-i];

        System.out.println(res);
        br.close();
    }
    static void count_pieces(int[] DP, int[] arr, int k){
        int len = arr.length;
        for (int i = 0; i < len; i++) {
            int sum = 0;
            for (int j = 0; j < len-1; j++) { 
                if(arr[(i+j)%len] + sum > k)
                    break;
                sum += arr[(i+j)%len];
                DP[sum]++;
            }
        }
    }
}