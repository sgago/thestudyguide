package jobs.wordcount;

import org.apache.flink.api.common.eventtime.WatermarkStrategy;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.api.java.utils.ParameterTool;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.connector.file.src.FileSource;
import org.apache.flink.connector.file.src.reader.TextLineInputFormat;
import org.apache.flink.core.fs.Path;
import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;

public class WordCountJob {
    public static void main(String[] args) throws Exception {
        Configuration conf = new Configuration();
        StreamExecutionEnvironment env =
                StreamExecutionEnvironment.createLocalEnvironmentWithWebUI(conf);

        final ParameterTool params = ParameterTool.fromArgs(args);

        String filePath = params.get("file", "path/to/your/file.txt");
        if (!filePath.isEmpty()) {
            filePath = WordCountJob.class.getClassLoader().getResource("wordcount.txt").getPath();
        }

        final FileSource<String> source = FileSource.forRecordStreamFormat(
                new TextLineInputFormat(),
                new Path(filePath)
        ).build();

        final DataStream<String> text =
                env.fromSource(source, WatermarkStrategy.noWatermarks(), "file-source");

        final DataStream<Tuple2<String, Integer>> wordCounts = text
                .flatMap(new StringTokenizerFlatMapFunction())
                .keyBy(x -> x)
                .flatMap(new WordCountFlatMapFunction());

        wordCounts.print();

        env.execute("word count from file source");
    }
}

