package jobs.simple;

import org.apache.flink.configuration.Configuration;
import org.apache.flink.table.api.*;
import org.apache.flink.table.api.Table;
import org.apache.flink.table.types.DataType;
import org.apache.flink.types.Row;

import java.util.Arrays;
import java.util.List;

public class TableExample
{
    public static void main(String[] args) throws Exception
    {
        Configuration conf = new Configuration();

        EnvironmentSettings settings = EnvironmentSettings
            .newInstance()
            .withConfiguration(conf)
            .inStreamingMode() // or .inBatchMode()
            .build();

        // Create a TableEnvironment for batch or streaming execution.
        // See the "Create a TableEnvironment" section for details.
        TableEnvironment tableEnv = TableEnvironment
            .create(settings);

        // Define in-memory data
        List<Row> people = Arrays.asList(
            Row.of("Alice", 12),
            Row.of("Bob", 10),
            Row.of("Carol", 18)
        );

        DataType rowDataType = DataTypes.ROW(
            DataTypes.FIELD("name", DataTypes.STRING()),
            DataTypes.FIELD("age", DataTypes.INT())
        );

        // Insert in-memory data into the source table
        Table sourceTable = tableEnv
            .fromValues(rowDataType, people);

        // Create a Table object from a Table API query
        tableEnv.createTemporaryView("SourceView", sourceTable);

        // Create a sink table with an explicitly defined schema
        tableEnv.executeSql(
        "CREATE TEMPORARY TABLE SinkTable (" +
                "name STRING, " +
                "age INT" +
                ") WITH ('connector' = 'print')"
        );

        // Create a Table object from a SQL query
        Table result = tableEnv
            .sqlQuery("SELECT name, age FROM SourceView WHERE age > 10");

        // Emit a Table API result Table to a TableSink, same for SQL result
        result.executeInsert("SinkTable").await();
    }
}
