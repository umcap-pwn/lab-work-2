import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        String input = scanner.nextLine();
        String[] emails = input.split(",\\s*");
        
        Set<String> uniqueEmails = new HashSet<>();
        
        for (String email : emails) {
            String normalized = normalizeEmail(email.trim());
            if (!normalized.isEmpty()) {
                uniqueEmails.add(normalized);
            }
        }
        
        System.out.println(uniqueEmails.size());
        scanner.close();
    }
    
    private static String normalizeEmail(String email) {
        int atIndex = email.indexOf('@');
        if (atIndex == -1) {
            return "";
        }
        
        String local = email.substring(0, atIndex);
        String domain = email.substring(atIndex);
        
        int plusIndex = local.indexOf('+');
        String usernameToCheck = local;
        if (plusIndex != -1) {
            usernameToCheck = local.substring(0, plusIndex);
        }
        
        if (!isValidUsername(usernameToCheck)) {
            return "";
        }
        
        StringBuilder normalizedLocal = new StringBuilder();
        
        for (int i = 0; i < local.length(); i++) {
            char c = local.charAt(i);
            if (c == '+') {
                break;
            }
            if (c != '.') {
                normalizedLocal.append(c);
            }
        }
        
        return normalizedLocal.toString().toLowerCase() + domain;
    }
    
    private static boolean isValidUsername(String username) {
        if (username.length() > 30) {
            return false;
        }
        
        if (username.length() > 0 && (username.charAt(0) == '.' || username.charAt(username.length() - 1) == '.')) {
            return false;
        }
        
        boolean hasConsecutiveDots = false;
        for (int i = 0; i < username.length(); i++) {
            char c = username.charAt(i);
            
            if (i > 0 && c == '.' && username.charAt(i - 1) == '.') {
                hasConsecutiveDots = true;
                break;
            }
            
            if (!isAllowedCharacter(c)) {
                return false;
            }
        }
        
        if (hasConsecutiveDots) {
            return false;
        }
        
        return true;
    }
    
    private static boolean isAllowedCharacter(char c) {
        if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
            return true;
        }
        if (c >= '0' && c <= '9') {
            return true;
        }
        if (c == '.') {
            return true;
        }
        
        return false;
    }
}