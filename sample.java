import com.sun.jna.*;

public class sample {

    public interface GoSO extends Library {
        GoSO INSTANCE = (GoSO) Native.loadLibrary("sample" ,GoSO.class);
        void Hello(String s);
        String Hello2(String s);
    }
    public static void main(String[] args) {
        GoSO.INSTANCE.Hello("World");
        System.out.println(GoSO.INSTANCE.Hello2("New World"));
    }
}