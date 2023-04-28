import org.junit.Test;
import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import static org.junit.Assert.*;

public class HelloActionsTest {
  @Test
  public void testHelloActions() {
    ByteArrayOutputStream outContent = new ByteArrayOutputStream();
    System.setOut(new PrintStream(outContent));
    HelloActions.main(new String[]{});
    assertEquals("Hello, GitHub Actions!\n", outContent.toString());
    System.setOut(System.out);
  }
}

