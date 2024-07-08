# 型アサーションと型ガード

## as による型アサーション

- 型アサーションとは
  - 開発者が「この値はこの型だ」と明示的に指定する方法
  - コンパイラに対して形を矯正するが、型安全性が保証されない可能性がある
  - 使用方法
    - ```typescript
      const value: any = 'hello';
      const length = (value as string).length;
      ```
- 型アサーションと型キャストの違い
  - 型キャストは値の変換を伴うが、型アサーションはコンパイラに型を信じ込ませるだけで実際の値の変換は行わない
  - 型アサーションは実行時に型チェックを行わないため、ランタイムエラーの原因となる可能性がある

## 型ガードでスマートに型安全を保証する

- 型ガードとは
  - あるスコープ内で特定の型を保証するためのチェックを行う式のこと
  - 例: `typeof` 演算子や `instanceof`　演算子を使用して、変数の型を確認する
  - 使用方法
    - ```typescript
      const value: unknown = "hello";
      if (typeof value === "string") {
        console.log(value.length); // valueはstring型として扱われる
      }
      ```
  - ユーザー定義の型ガード:
    - 型述語（Type Predicate）を使って独自の型ガードを定義
    - arg is Typeの形で定義し、特定の条件を満たすかどうかをチェック
    - 例:
      - ```typescript
        type User = { username: string; address: { zipcode: string; town: string } };
        const isUser = (arg: unknown): arg is User => {
          const u = arg as User;
          return (
            typeof u.username === "string" &&
            typeof u.address === "object" &&
            typeof u.address.zipcode === "string" &&
            typeof u.address.town === "string"
          );
        };
        ```
  - ユーザー定義の型ガードの使用例:
    - 型ガード関数を使って、特定の型かどうかを判定し、安全にプロパティにアクセスする
    - 例:
      - ```typescript
        const data: unknown = JSON.parse(`{ "username": "patty", "address": { "zipcode": "12345", "town": "Maple Town" } }`);
        if (isUser(data)) {
          console.log(data.username); // "patty"
          console.log(data.address.town); // "Maple Town"
        } else {
          console.error('Invalid user data');
        }
        ```

## まとめ
- 型アサーションは、開発者が特定の型を強制的に指定する手段であり、型安全性を保証しない
- 型ガードは、特定の条件を満たすかどうかをチェックし、型安全性を確保する方法
- arg is Type構文を使うことで、TypeScriptがその型を保証するようにできる