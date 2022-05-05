import com.google.common.base.Charsets;
import com.google.common.base.Supplier;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.Map;
import java.util.NoSuchElementException;
import java.util.Random;

public class Main {
    private static String[] nounList = {
        "cloud",
        "elephant",
        "gø language",
        "laptøp",
        "cøntainer",
        "micrø-service",
        "turtle",
        "whale",
        "gøpher",
        "møby døck",
        "server",
        "bicycle",
        "viking",
        "mermaid",
        "fjørd",
        "legø",
        "flødebolle",
        "smørrebrød"
    };

    private static String[] verbList = {
        "will drink",
        "smashes",
        "smøkes",
        "eats",
        "walks tøwards",
        "løves",
        "helps",
        "pushes",
        "debugs",
        "invites",
        "hides",
        "will ship"
    };

    private static String[] adjectivesList = {
        "the exquisite",
        "a pink",
        "the røtten",
        "a red",
        "the serverless",
        "a brøken",
        "a shiny",
        "the pretty",
        "the impressive",
        "an awesøme",
        "the famøus",
        "a gigantic",
        "the gløriøus",
        "the nørdic",
        "the welcøming",
        "the deliciøus"
    };

    private static Map<String, String[]> wordList = Map.of("nouns", nounList, "verbs", verbList, "adjectives", adjectivesList);

    public static void main(String[] args) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);
        server.createContext("/noun", handler(() -> randomWord("nouns")));
        server.createContext("/verb", handler(() -> randomWord("verbs")));
        server.createContext("/adjective", handler(() -> randomWord("adjectives")));
        server.start();
    }

    private static String randomWord(String table) {
        if(!wordList.containsKey(table)) {
            throw new NoSuchElementException(table);
        }

        Random r = new Random();
        int randomNumber = r.nextInt(wordList.get(table).length);
        return wordList.get(table)[randomNumber];
    }

    private static HttpHandler handler(Supplier<String> word) {
        return t -> {
            String response = "{\"word\":\"" + word.get() + "\"}";
            byte[] bytes = response.getBytes(Charsets.UTF_8);

            System.out.println(response);
            
            t.getResponseHeaders().add("content-type", "application/json; charset=utf-8");
            t.getResponseHeaders().add("cache-control", "private, no-cache, no-store, must-revalidate, max-age=0");
            t.getResponseHeaders().add("pragma", "no-cache");

            t.sendResponseHeaders(200, bytes.length);

            try (OutputStream os = t.getResponseBody()) {
                os.write(bytes);
            }
        };
    }
}
