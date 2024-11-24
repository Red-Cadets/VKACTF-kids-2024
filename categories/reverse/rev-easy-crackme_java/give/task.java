import java.util.*;

class Check {
    public static void main(String args[]) {
        Check Check = new Check();
        Scanner scanner = new Scanner(System.in); 
        System.out.print("Enter check password: ");
        String userInput = scanner.next();
        if (userInput.startsWith("vka{") && userInput.endsWith("}")) {
            String input = userInput.substring("vka{".length(), userInput.length() - 1);
            if (Check.checkPassword(input)) {
                System.out.println("Good job!!!");
            } else {
                System.out.println("Try harder");
            }
        } else {
            System.out.println("Nope :(");
        }
    }
    public boolean checkPassword(String password) { 
        boolean part1 = password.startsWith("jaVA_CoD3_");
        boolean part3 = password.endsWith("5y_TO_ReAd");
        boolean part2 = password.contains("Is_tHa7_e4");        
        return part1 && part3 && part2;
    }
}
