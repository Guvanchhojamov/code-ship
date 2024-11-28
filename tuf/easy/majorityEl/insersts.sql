select * from update_device_token_client(12,'533A109D-594B-4C8A-A128-CE0B71FC88FF','iOS','76cdda86d9c70183b47fa25eaf1aec921fa1fe8a95e9541de5209b7c34610330
', now()::timestamp);

select * from update_device_token_client(12,'A7CCC26A-16AC-404E-A671-E02C4AA71C44','iOS','0a84396d19436a7ec4c70e4a89569d7399ec0d0d55de005123ab5bee99c76542
', now()::timestamp);


--- 

select * from update_device_token_client(10679,'5342FDE1-C358-41D4-AF9B-F870DDEC38A7','iOS','7411fd9509d0b8367d315dd11ba75a50592d450de2f8c4d54c70d6878de32a68', now()::timestamp);
select * from update_device_token_client(10679,'F758B420-3DA7-4197-9A13-CAEBC7B47A85','iOS','8933797ba7bb716d7c0d3393dd8af5800816822a9da3b76b21d192be68dedeaf', now()::timestamp);
select * from update_device_token_client(10679,'A732F79B-D7FF-4596-B460-9AAA86923F3A','iOS','9c6bea478774e899f19beae87b1943f04c1c495454e1613d5f71f280d37bd9fc', now()::timestamp);
select * from update_device_token_client(nil,'CC49B916-9D65-4632-BD2F-56AF9420EF32','iOS','74e904627d17fdf3d627e90ebc4a104f7714d58e123ce1a2e330fa93661132c3', now()::timestamp);


--- 


curl -v \
 -X PUT \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJidlBheWxvYWQiOnsidXNlcklkIjoxMDY3OSwiZGV2aWNlSWQiOiI1NTZiMzc2YS1kOWZlLTQwNTMtOTdjYS03MTU5ZWIxNjc0MmQifSwiaWF0IjoxNzMxMjQ0MjgwfQ.c72hHOkEEtRlUzKcpFX7toUFVvYsVSb5KfvecRHboUE" \
 -H "Accept-Language: en" \
 -H "BV-Token: hackMeIfYouCan" \
 -H "RestrictedMode: false" \
 -H "platformType: iOS" \
 -H "BV-Device-Id: 73EA8949-FC61-41DA-8E0D-3D1B47E04A71" \
 -H "BV-Device-Token: a516ac29c5f075e6a6cfa2f0702ec77cde7d1e6e2fb5d0a25a8ab528fef814f4" \
 -H "Content-Type: application/json" \
 -H "versionCode: 47" \
 "https://gateway.beletvideo.com/v1/device/updateToken"