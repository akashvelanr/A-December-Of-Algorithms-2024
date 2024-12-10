import java.util.ArrayList;
import java.util.List;
import java.util.*;

public class Main{
    public static void main(String[] args) {
        Scanner scn = new Scanner(System.in);
        System.out.println("Enter no. of floors:");
        int n = scn.nextInt();
        List<List<Integer>> tower=new ArrayList<>(
        				Arrays.asList(
        												Arrays.asList(1),Arrays.asList(1,1)
       					 )); 
        //System.out.println(tower);
        for(int i=2;i<n;i++){
            List<Integer> f=new ArrayList<>();
            f.add(1);
            List<Integer> g=tower.get(i-1);
            for(int j=0;j<i-1;j++){
                f.add(g.get(j)+g.get(j+1));
            }
            f.add(1);
            tower.add(f);
        }
        System.out.println("Output:"+tower.subList(0,n));
    }
}