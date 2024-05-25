package jobs.wordcount;

import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.util.Collector;

public class StringTokenizerFlatMapFunction implements FlatMapFunction<String, String> {
    @Override
    public void flatMap(@org.jetbrains.annotations.NotNull String s, Collector<String> collector) throws Exception {
        String[] words = s.toLowerCase().split("\\s+");

        for (String word : words) {
            if (!word.isEmpty()) {
                collector.collect(word);
            }
        }
    }
}
