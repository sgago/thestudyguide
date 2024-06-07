package jobs.simple;

import org.apache.flink.configuration.Configuration;
import org.apache.flink.connector.datagen.table.DataGenConnectorOptions;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.table.api.*;
import org.apache.flink.table.api.Table;
import org.apache.flink.table.types.DataType;
import org.apache.flink.types.Row;

import java.util.Arrays;
import java.util.List;

public class SimpleTable
{
    public static void main(String[] args) throws Exception
    {
        Configuration conf = new Configuration();

        StreamExecutionEnvironment env = StreamExecutionEnvironment
            .createLocalEnvironmentWithWebUI(conf);

        EnvironmentSettings settings = EnvironmentSettings
            .newInstance()
            .withConfiguration(conf)
            .inStreamingMode()
//                .inBatchMode()
            .build();

        // Create a TableEnvironment for batch or streaming execution.
        // See the "Create a TableEnvironment" section for details.
        TableEnvironment tableEnv = TableEnvironment.create(settings);

        // Define in-memory data
        List<Row> data = Arrays.asList(
            Row.of("Alice", 12),
            Row.of("Bob", 10),
            Row.of("Carol", 18)
        );

        Schema schema = Schema.newBuilder()
            .column("name", DataTypes.STRING())
            .column("age", DataTypes.INT())
            .build();

        TableDescriptor sourceDescriptor = TableDescriptor
            .forConnector("datagen")
            .schema(schema)
            .option(DataGenConnectorOptions.ROWS_PER_SECOND, 100L)
            .build();

        DataType rowDataType = DataTypes.ROW(
            DataTypes.FIELD("name", DataTypes.STRING()),
            DataTypes.FIELD("age", DataTypes.INT())
        );

        // Insert in-memory data into the source table
        org.apache.flink.table.api.Table sourceTable = tableEnv
            .fromValues(rowDataType, data);

        // Create a Table object from a Table API query
        tableEnv.createTemporaryView("SourceTable", sourceTable);

        // Create a sink table (using SQL DDL)
        tableEnv.executeSql(
                "CREATE TEMPORARY TABLE SinkTable " +
                        "WITH ('connector' = 'print') " +
                        "LIKE SourceTable (EXCLUDING OPTIONS)");

        // Create a Table object from a SQL query
        Table result = tableEnv
            .sqlQuery("SELECT name, age FROM SourceTable WHERE age > 10");

        // Emit a Table API result Table to a TableSink, same for SQL result
        result.executeInsert("SinkTable").await();


    }
}
