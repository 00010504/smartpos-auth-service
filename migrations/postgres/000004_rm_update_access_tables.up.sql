INSERT INTO "permission" ("id","section_id" ,"name", "translation", "created_by") VALUES 
    -- catalog
    ('2a1c1dc3-936d-423b-aae6-7037dee93f48', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('46ff423e-d7f3-478c-9d4d-57959ee18cc2', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('389c4d7e-6f33-4e60-9092-b45cb649b074', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('48463d76-df0e-4434-ba65-407e616b77dc', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'print labels', '{"us":"Yorliq chiqarish", "en":"Print labels", "ru":"Распечатать ценники"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('34f7535b-242c-467d-8522-6831b2bb1fc1', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'see purchase costs of the products', '{"us":"Mahsulotlarni sotib olish narhlarini ko''rish", "en":"See purchase costs of the products", "ru":"Посмотреть стоимость покупки продуктов"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('88d11654-e087-4551-a546-3f0b973d153b', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'can change price', '{"us":"Narxni o''zgartirishi mumkin", "en":"Can chage price", "ru":"Можно изменить цену"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('6b615cb4-5137-4750-988d-f9919c279fed', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'can change product stock', '{"us":"Mahsulot zaxirasini o''zgartirishi mumkin", "en":"Can change product stock", "ru":"Может изменить запас продукта"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('075fa6b0-f281-4179-ae5c-a2b25a56cf19', '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),

-- writeOff
    ('6b4bb3f2-bd24-4863-bec8-9e065aed1121', '71967e9c-3a16-410c-82f0-608ff4bbe7cf', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('d4a428dd-ecd7-427b-8990-a77fc5166b6d', '71967e9c-3a16-410c-82f0-608ff4bbe7cf', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('62925f23-db06-4ede-8b55-2e579c6e0f44', '71967e9c-3a16-410c-82f0-608ff4bbe7cf', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('1f1a647e-f060-4acc-a97b-944a124e1484', '71967e9c-3a16-410c-82f0-608ff4bbe7cf', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),

    -- import
    ('d25eec66-153d-485d-84da-4e039264889d', '2dce3c0f-123c-4b5e-af52-e267437787f6', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('940e7638-49f8-4051-897f-4357cbbc7b7b', '2dce3c0f-123c-4b5e-af52-e267437787f6', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('14866bdd-a498-4a2f-940a-873d32b35441', '2dce3c0f-123c-4b5e-af52-e267437787f6', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('ef8c93bf-cff4-441b-9c3f-d607c7ba4f69', '2dce3c0f-123c-4b5e-af52-e267437787f6', 'can see sales of the import', '{"us":"Import sotuvini ko''rish mumkin", "en":"Can see sales of the import", "ru":"Можно увидеть продажи импорта"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('db58c3ed-bdce-4a36-b19c-4b520a547596', '2dce3c0f-123c-4b5e-af52-e267437787f6', 'can see costs of the product', '{"us":"Mahsulot narxini ko''rish mumkin", "en":"Can see Costs of the product", "ru":"Можно увидеть стоимость продукта"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- transfer
    ('1d9def8a-d940-412c-982b-0aaeeb995b9a', '23fc53a1-3a37-41c1-9ad0-e985067cb0ba',  'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('62d8a3f7-8981-4b3f-95e1-de62fbb52bb3', '23fc53a1-3a37-41c1-9ad0-e985067cb0ba',  'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('dbcbeada-6087-4e2e-8a03-79b41789a5bf', '23fc53a1-3a37-41c1-9ad0-e985067cb0ba',  'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('c4921cfd-9274-47b3-9d5c-1dc0cc07a4a2', '23fc53a1-3a37-41c1-9ad0-e985067cb0ba',  'accept and verify transfers', '{"us":"O''tkazmani qabul qiling va tasdiqlang", "en":"Accept and verify transfer", "ru":"Принять и подтвердить перевод"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),

      -- orders
    ('870e957a-2659-4afd-ac4b-a5962de50d3b', '54e0d4a1-2df2-42c8-90c3-efdf5cce74bd', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('12ccb7b9-42f5-4d90-b320-569954dbd247', '54e0d4a1-2df2-42c8-90c3-efdf5cce74bd', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('759c615e-2988-46da-aa1c-bfbf50a4b170', '54e0d4a1-2df2-42c8-90c3-efdf5cce74bd', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),

    -- repricing
    ('0a60dc1f-b704-42bc-a17b-b686a5a55a97', '19d97f84-112f-4d8a-a3e6-787f3c218acf', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('06a0c2ca-5fba-41c9-9e3f-62f65759ca6a', '19d97f84-112f-4d8a-a3e6-787f3c218acf', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('e39716d6-7c24-4717-b67e-274b84884bf0', '19d97f84-112f-4d8a-a3e6-787f3c218acf', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('71a6257f-ca0e-49fb-a38a-68e52c9abbd0', '19d97f84-112f-4d8a-a3e6-787f3c218acf', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),

     -- labels
    ('a4a4041b-79f2-4b8e-89f4-ea0639602c04', '8243efd3-58da-4c1a-a7c3-e85d3d1cc51d', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('a56825aa-d5b1-4356-b9b2-2b2651c021a1', '8243efd3-58da-4c1a-a7c3-e85d3d1cc51d', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('3c28c452-9508-4b07-b86f-633eb7de7915', '8243efd3-58da-4c1a-a7c3-e85d3d1cc51d', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('8e1ba694-c649-4528-a97b-dff3d3de3860', '8243efd3-58da-4c1a-a7c3-e85d3d1cc51d', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),

    -- suppliers
    ('8a373b5d-2dea-436f-96e4-b5990fa72bea', '6175bc30-fc74-4546-aa28-98859e4935c1', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('02ed7f0a-d4be-4a70-bfd0-84ce6807594d', '6175bc30-fc74-4546-aa28-98859e4935c1', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('82b2ba44-fe8f-4331-90fa-21cae157b98e', '6175bc30-fc74-4546-aa28-98859e4935c1', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('cd17eefa-a7e9-4678-bf37-bedfe2d72186', '6175bc30-fc74-4546-aa28-98859e4935c1', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- -- categories
    ('482b5933-b2e5-43c0-81d3-e55cbbcaf3ca', 'ce52c92f-2a8b-430c-8dd7-1093810ed5c5', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('046d9eeb-a9c4-421e-8036-04ba7fb13f88', 'ce52c92f-2a8b-430c-8dd7-1093810ed5c5', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('5dcf7d96-82f2-4b8e-9bdb-35d0af37b54f', 'ce52c92f-2a8b-430c-8dd7-1093810ed5c5', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('98770c78-4016-4170-8e06-a9c4dd820110', 'ce52c92f-2a8b-430c-8dd7-1093810ed5c5', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- new sale create sales
    ('53c360f9-a10b-4f4b-8b52-670c3b8f87ca', 'a4c673a9-0b0c-4689-9feb-bf954fec2e06', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('d45d234a-e2eb-4861-b825-37217f821142', 'a4c673a9-0b0c-4689-9feb-bf954fec2e06', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('888fde2d-812c-4e5e-9e51-601e367eecbf', 'a4c673a9-0b0c-4689-9feb-bf954fec2e06', 'apply manual discounts', '{"us":"Qo''lda chegirmalarni qo''llang", "en":"Apply manual discounts", "ru":"Применение ручных скидок"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('e84a7e7d-2c6e-441f-82b7-1eeaf3a361fe', 'a4c673a9-0b0c-4689-9feb-bf954fec2e06', 'refund product', '{"us":"Mahsulotni qaytarish", "en":"Refund product", "ru":"Возврат продукта"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('f89ca40a-f2f9-4745-bebc-9cf293521c95', 'a4c673a9-0b0c-4689-9feb-bf954fec2e06', 'sale as debt', '{"us":"Qarz sifatida sotish", "en":"Sale as Debt", "ru":"Продажа в качестве долга"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('49f674fc-efbf-42fc-bb7c-2904c9c15a07', 'a4c673a9-0b0c-4689-9feb-bf954fec2e06', 'stock', '{"us":"Aksiya", "en":"Stock", "ru":"Акция"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- all sales
    ('8f5f5316-becc-44b9-9c50-ed91ef430733', '8921a4ef-23b3-4608-93bb-54a16907a2e0', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('4d8c05a0-fa2b-468a-9c87-f5adc07cee23', '8921a4ef-23b3-4608-93bb-54a16907a2e0', 'print reports', '{"us":"Hisobotlarni chop etish", "en":"Print reports", "ru":"Печать отчетов"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('25ae23a7-d395-4f99-afb4-4fcec02ed1b6', '8921a4ef-23b3-4608-93bb-54a16907a2e0', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('d24b0e6c-cc77-4e5b-935b-25a65a98f720', '8921a4ef-23b3-4608-93bb-54a16907a2e0', 'receipts history', '{"us":"Qabul qilish tarixi", "en":"Receipts history", "ru":"История чеков"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('6f79f732-777c-4c76-98d9-d519ebf7ec78', '8921a4ef-23b3-4608-93bb-54a16907a2e0', 'create new customer', '{"us":"Yangi mijoz yarating", "en":"Create new Customer", "ru":"Создать нового клиента"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- shifts open shift
    ('1a9b575b-31f5-47b7-9056-8875889921cd', '1202a5b0-dbc1-45f3-b774-84c5670eb7b1', 'open shift', '{"us":"Ochiq smena", "en":"Open shift", "ru":"Открытая смена"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('e7ffe144-757b-420e-b57d-4f04095b97e9', '1202a5b0-dbc1-45f3-b774-84c5670eb7b1', 'X reports', '{"us":"X xabar berish", "en":"X reports", "ru":"Х отчеты"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('2b641aa6-26b1-4d5b-9fc4-f5538338aec2', '1202a5b0-dbc1-45f3-b774-84c5670eb7b1', 'Y reports', '{"us":"Y xabar berish", "en":"Y reports", "ru":"Y отчеты"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- operations
    ('35289f03-a794-4013-8cda-2705545d969b', '445b8f2d-b13f-40d8-9946-489ca936926a', 'create incomes', '{"us":"Daromad yaratish", "en":"Create Incomes", "ru":"Создание доходов"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('7a9a4427-7d42-428c-b60e-4dd574e688cf', '445b8f2d-b13f-40d8-9946-489ca936926a', 'create expenses', '{"us":"Xarajatlarni yaratish", "en":"Create Expenses", "ru":"Создать расходы"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('d0848be0-0fa2-461d-a666-d7473c4c38cc', '445b8f2d-b13f-40d8-9946-489ca936926a', 'view reports', '{"us":"Hisobotlarni ko''rish", "en":"View reports", "ru":"Просмотр отчетов"}', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- employees
    ('06247d7f-b162-45bf-b447-920fd5b6c544', '93ebc9a0-9400-4864-af1c-0f0acaab943f', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('4aa93cae-fa63-4b20-a2fa-427df64cb5f6', '93ebc9a0-9400-4864-af1c-0f0acaab943f', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('67348e49-5cf8-42d9-a5cc-82fae9057498', '93ebc9a0-9400-4864-af1c-0f0acaab943f', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('fada9b62-b7ed-4705-8749-c9974d451273', '93ebc9a0-9400-4864-af1c-0f0acaab943f', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- roles view only
    ('e57a1487-380e-4e40-a114-b3f62fb24487', '80816c70-ca54-4cd8-b8a7-4733799406f0', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('894e5409-1268-424c-a027-0240212caa09', '80816c70-ca54-4cd8-b8a7-4733799406f0', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('82bdce7c-4f90-4365-bfaa-3145d061634f', '80816c70-ca54-4cd8-b8a7-4733799406f0', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('3936b622-c5c8-4b56-a2e2-a21b4f1dba24', '80816c70-ca54-4cd8-b8a7-4733799406f0', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- hr
    -- time cards
    -- all clients
    ('9a3644e8-2b3f-46a7-afce-0280e457e502', '313db387-84c4-4365-99b0-a01ac85d12a0', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('f048dc97-5a56-4eec-8cb6-f929e45f3674', '313db387-84c4-4365-99b0-a01ac85d12a0', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('d0e2acb3-81b9-4f7f-993d-94959467eecc', '313db387-84c4-4365-99b0-a01ac85d12a0', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('a430f2fc-d390-4448-b66f-24760e70ef64', '313db387-84c4-4365-99b0-a01ac85d12a0', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- clients debts
    -- loyalty program
    -- client groups
    -- settings -> profile
    ('3540610d-eb77-45b0-b78a-ed3ee3bfeefe', '32f5e61a-8f63-48ed-9c0f-9449ac5f74dc', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- settings -> company
    ('63685ec6-72b2-41c4-a992-3e281397004e', '63685ec6-72b2-41c4-a992-3e281397004e', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- settings -> cashbox
    ('f519f52d-c0b0-4b84-887e-f3f2ec0957be', '9541320f-f8bc-41ab-bc63-2dcc9397fb69', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('fd281359-2798-4869-84af-a5324b73bc0e', '9541320f-f8bc-41ab-bc63-2dcc9397fb69', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('fa2a1dad-c779-4112-af0a-3b2391548ecf', '9541320f-f8bc-41ab-bc63-2dcc9397fb69', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('a0bb5ff4-23c2-4802-904c-bde6a8fc7abe', '9541320f-f8bc-41ab-bc63-2dcc9397fb69', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- settings -> products
    ('05fd5a57-f648-498a-aea9-02b30f788f43', '5c8ae38d-98c9-4128-a157-677ad3b7c858', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('5d0db786-8d9a-4d73-a657-bbf69791a6c8', '5c8ae38d-98c9-4128-a157-677ad3b7c858', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('101e3741-a81c-4989-907c-379e43035707', '5c8ae38d-98c9-4128-a157-677ad3b7c858', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('1cf1c912-2f4a-4d02-a93f-06354612d2e9', '5c8ae38d-98c9-4128-a157-677ad3b7c858', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- setings -> store
    ('c5260254-6945-43aa-8433-2fc8330edf8c', '66a2ff50-52e1-499f-9c8a-5ffc1082e79f', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('5124e71e-b283-4e96-9ed4-b22cda3e8bfc', '66a2ff50-52e1-499f-9c8a-5ffc1082e79f', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('6da81b69-2563-41f0-aadb-a8aeadd082a6', '66a2ff50-52e1-499f-9c8a-5ffc1082e79f', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('9de4ad71-53f1-4939-a1a4-f89cfcd18b2b', '66a2ff50-52e1-499f-9c8a-5ffc1082e79f', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- settings -> reciept
    ('e9fe92f2-751e-4798-9409-2e5f4911b558', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('9f4ae1e5-022c-4df2-b9f3-60c5912e818d', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('e9cdcc0d-bf8d-46f8-83bc-4143edc926d8', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('22de842e-8943-45b8-9656-16acf3aa6f30', '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- settings -> label
    ('07fce546-06f3-4fd3-8a3d-7a2b2dbd0f63', '02662be0-c2c9-4291-ae0d-982d89e0f20b', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('9bcfcef2-8f2a-4162-8678-6fcd1e7dd3a2', '02662be0-c2c9-4291-ae0d-982d89e0f20b', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('d45147ee-b9a6-4c4c-940c-606b1478058f', '02662be0-c2c9-4291-ae0d-982d89e0f20b', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('8ccd4e38-9a2b-4568-ac06-4860d8977b48', '02662be0-c2c9-4291-ae0d-982d89e0f20b', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    -- settings -> vat
    ('4247dec8-6b2c-402f-8e25-47383ec81fbc', '4f7ac098-95a0-4ad8-9c5e-135772144002', 'view only', '{"us":"Faqat ko''rish", "en":"View only", "ru":"Для просмотра"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('6a133e13-8412-4f04-8f3c-40055fe49fb9', '4f7ac098-95a0-4ad8-9c5e-135772144002', 'update', '{"us":"Yangilash", "en":"Update", "ru":"Обновить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('ca7d6391-71bc-403e-b817-cbaaf1ac4876', '4f7ac098-95a0-4ad8-9c5e-135772144002', 'create', '{"us":"Yaratish", "en":"Create", "ru":"Создать"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'),
    ('7aba659b-5494-4bce-baa0-9329393d845f', '4f7ac098-95a0-4ad8-9c5e-135772144002', 'delete', '{"us":"O''chirish", "en":"Delete", "ru":"Удалить"}','9a2aa8fe-806e-44d7-8c9d-575fa67ebefd')
    -- settings -> application
    -- settings -> payment
    -- settings -> storage
    -- settings -> profile
    -- settings -> kanban
    -- settings -> notification