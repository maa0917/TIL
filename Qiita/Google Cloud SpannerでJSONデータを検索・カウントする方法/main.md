# Google Cloud SpannerでJSON配列データを検索・カウントする方法

Google Cloud Spannerは、スケーラブルで高可用性を持つマネージドリレーショナルデータベースサービスです。
SpannerはJSONデータ型をサポートおり、これにより複雑なデータ構造を効率的に保存・操作することが可能です。
本記事では、JSON配列データの格納、検索、およびカウントの方法について説明します。

## テーブルの作成

まず、JSONデータを格納するためのテーブルを作成します。

```sql
CREATE TABLE example (
    id INT64 NOT NULL,
    data JSON
) PRIMARY KEY(id);
```

## データの挿入

次に、JSONデータをテーブルに挿入します。ここでは、名前、電話番号、住所を含むJSONデータを挿入します。

```sql
INSERT INTO example (id, data) VALUES (1, JSON '[{"name": "Alice", "phone": "123-4567", "address": "Tokyo"}, {"name": "Bob", "phone": "987-6543", "address": "Tokyo"}]');
INSERT INTO example (id, data) VALUES (2, JSON '[{"name": "Charlie", "phone": "555-5555", "address": "Osaka"},{"name": "Dave", "phone": "123-0000", "address": "Osaka"}]');
```

データの内容は次のようになります：

| id | data                                                                                                                      |
|----|---------------------------------------------------------------------------------------------------------------------------|
| 1  | [{"name": "Alice", "phone": "123-4567", "address": "Tokyo"}, {"name": "Bob", "phone": "987-6543", "address": "Tokyo"}]    |
| 2  | [{"name": "Charlie", "phone": "555-5555", "address": "Osaka"}, {"name": "Dave", "phone": "123-0000", "address": "Osaka"}] |

## JSONデータを検索し、カウント

特定の条件に一致するJSONデータを検索し、カウントするクエリを実行します。

```sql
SELECT COUNT(*) AS match_count
FROM example e
WHERE EXISTS (
    SELECT 1
    FROM UNNEST(JSON_QUERY_ARRAY(e.data)) AS json_array_element
    WHERE JSON_VALUE(json_array_element, '$.phone') = '123-4567'
      AND JSON_VALUE(json_array_element, '$.name') = 'Alice'
);
```

### JSON_QUERY_ARRAYについて

JSON_QUERY_ARRAYは、JSONデータの配列要素を抽出するための関数です。
この関数を使用すると、JSON配列を個々の要素に分解し、それぞれに対して操作を行うことができます。
例のクエリでは、dataフィールドからJSON配列を展開し、各要素を行として取得します。

### UNNESTについて

UNNEST関数は、配列やリストのようなコレクションを個々の要素に分解して、各要素を行として展開します。
これにより、配列の各要素に対してクエリを実行できます。
例のクエリでは、UNNESTを使用してJSON配列を展開し、各要素を取得します。

### JSON_VALUEについて

JSON_VALUEは、JSONオブジェクトからスカラー値（文字列、数値など）を抽出する関数です。
JSONパスを指定して、特定のキーに対応する値を取得します。
例のクエリでは、JSON配列の各要素から電話番号と名前を取得します。

### EXISTSについて

EXISTSを使用することで、サブクエリ内で条件に一致するレコードが存在するかどうかを効率的に確認できます。
例のクエリでは、各exampleレコードに対して、dataフィールド内のJSON配列から特定の電話番号と名前を持つ要素が存在するかどうかをチェックします。
条件を満たす最初のレコードが見つかるとサブクエリの処理が終了するため、効率的です。
`SELECT 1` で具体的な値（ここでは1）は重要ではなく、単に条件を満たす行が存在するかどうかを確認するために使用されます。

## まとめ

この記事では、Google Cloud Spannerを使用してJSON配列データの検索クエリの実行方法を解説しました。
また、JSON_QUERY_ARRAY、UNNEST、JSON_VALUE、およびEXISTSの各関数について説明しました。
これらの機能を活用することで、複雑なデータ構造を簡単に操作し、データベースの柔軟性とパフォーマンスを向上させることができます。
ぜひ試してみてください。

## 参考

- [JSON functions in GoogleSQL](https://cloud.google.com/spanner/docs/reference/standard-sql/json_functions)