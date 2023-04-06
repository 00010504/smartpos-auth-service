CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TYPE user_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', 
    '9fb3ada6-a73b-4b81-9295-5c1605e54552' 
);

CREATE TYPE app_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9fb3ada6-a73b-4b81-9295-5c1605e54552' 
);

CREATE TABLE "user_status" (
    "id" UUID PRIMARY KEY, 
    "name" VARCHAR NOT NULL, 
    "translation" JSONB NOT NULL,
    UNIQUE("name")
);

INSERT INTO "user_status" ("id", "name", "translation") VALUES
    ('1fe92aa8-2a61-4bf1-b907-182b497584ad','invited', '{"en":"Invited", "ru":"Приглашенный", "uz":"Taklif qilingan"}'), -- pending
    ('9fb3ada6-a73b-4b81-9295-5c1605e54552','active', '{"en":"Active", "ru":"Активный", "uz":"Faol"}'), -- active
    ('0adc982c-749b-4446-9d36-d136a76b99ea', 'rejected','{"en":"Rejected", "ru":"Отклоненный", "uz":"Rad etilgan"}'), -- rejected
    ('3e6eff54-dd23-4603-99f6-8f5fc24d19ff', 'blocked','{"en":"Blocked", "ru":"Заблокирован", "uz":"Bloklangan"}'); -- blocked

CREATE TABLE IF NOT EXISTS ui_mode (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(200) NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE ("name")
);

INSERT INTO "ui_mode" (
    "id",
    "name",
    "translation"
) VALUES
    ('0adc982c-749b-4446-9d36-d136a76b99ea', 'dark',  '{"en":"Dark", "ru":"Темный", "uz":"Tungi"}' ), 
    ('3e6eff54-dd23-4603-99f6-8f5fc24d19ff', 'light', '{"en":"Light", "ru":"Яркий", "uz":"Kunduzgi"}'),
    ('dc4f3251-a573-4eb4-868e-dfe0bb6aafa3', 'system', '{"en":"System", "ru":"Система", "uz":"Tizim"}');


CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY,
    "phone_number" VARCHAR(20) NOT NULL,
    "is_validated" BOOLEAN NOT NULL DEFAULT FALSE,
    "user_type_id" user_type NOT NULL DEFAULT '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    "last_version" INT NOT NULL DEFAULT 0,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    UNIQUE ("phone_number", "deleted_at", "is_validated")
);
CREATE INDEX IF NOT EXISTS "user_deleted_at_idx" ON "user"("deleted_at");


-- admin insert
INSERT INTO "user" (
    "id",
    "phone_number",
    "user_type_id",
    "is_validated",
    "created_by"
) VALUES (
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '99894172774',
    '9fb3ada6-a73b-4b81-9295-5c1605e54552',
    TRUE,
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'
);

CREATE TABLE IF NOT EXISTS "company" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(300) NOT NULL,
    "owner" UUID NOT NULL,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID,
    UNIQUE ("name", "deleted_at", "created_by")
);


CREATE INDEX IF NOT EXISTS "company_deleted_at_idx" ON "company"("deleted_at");


CREATE TABLE IF NOT EXISTS "company_user" (
    "id" UUID PRIMARY KEY,
    "external_id" VARCHAR NOT NULL,
    "status" UUID NOT NULL DEFAULT '9fb3ada6-a73b-4b81-9295-5c1605e54552' REFERENCES "user_status"("id") ON DELETE CASCADE,
    "company_id" UUID,
    "user_id" UUID NOT NULL REFERENCES "user"("id"),
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    UNIQUE ("company_id", "user_id", "deleted_at")
);

CREATE INDEX IF NOT EXISTS "company_user_deleted_at_idx" ON "company_user"("deleted_at");


CREATE TABLE IF NOT EXISTS "shop" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(300) NOT NULL,
    "company_id" UUID,
    "created_by" UUID,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID,
    UNIQUE ("name", "deleted_at", "company_id")
);

CREATE INDEX "shop_deleted_at_idx" ON "shop"("deleted_at");

CREATE TABLE IF NOT EXISTS "shop_user" (
    "shop_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    UNIQUE ("shop_id", "user_id", "deleted_at")
);
CREATE INDEX IF NOT EXISTS "shop_user_deleted_at_idx" ON "shop_user"("deleted_at");


CREATE TABLE IF NOT EXISTS "user_profile" (
    "id" UUID PRIMARY KEY, 
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "image" TEXT NOT NULL DEFAULT '',
    "password" TEXT NOT NULL,
    "last_company_id" UUID,
    "mode_id" UUID NOT NULL DEFAULT '3e6eff54-dd23-4603-99f6-8f5fc24d19ff' REFERENCES "ui_mode"("id"),
    "user_id" UUID NOT NULL REFERENCES "user"("id") ON DELETE CASCADE,
    "version" INT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE ("user_id", "version")
);


CREATE TABLE IF NOT EXISTS "role" ( 
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "is_deletable" BOOLEAN NOT NULL DEFAULT TRUE,
    "company_id" UUID,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    UNIQUE ("name", "company_id", "deleted_at")
); 

CREATE INDEX "role_deleted_at_idx" ON "role"("id");

CREATE TABLE "user_role" (
    "role_id" UUID NOT NULL,
    "company_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    PRIMARY KEY ("role_id", "company_id", "user_id", "deleted_at")
);

CREATE TABLE "module" ( 
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "app_type_id" app_type NOT NULL,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "translation" JSONB NOT NULL,
    UNIQUE("name", "deleted_at")
); 

CREATE TABLE IF NOT EXISTS "section" ( 
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    "module_id"  UUID NOT NULL REFERENCES "module"("id"),
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS "section_deleted_at_idx" ON "section"("deleted_at"); 

CREATE TABLE IF NOT EXISTS "permission" ( 
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    "section_id" UUID NOT NULL REFERENCES "section"("id") ON DELETE CASCADE,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    UNIQUE ("name", "deleted_at", "section_id")
); 

CREATE INDEX IF NOT EXISTS "permission_deleted_at_idx" ON "permission"("deleted_at");


CREATE TABLE IF NOT EXISTS "permission_role" (
    "role_id" UUID REFERENCES "role"("id") ON DELETE  SET NULL,
    "permission_id" UUID REFERENCES "permission"("id") NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID REFERENCES "user"("id") ON DELETE SET NULL,
    UNIQUE ("role_id", "permission_id", "deleted_at")
);

CREATE INDEX IF NOT EXISTS "permission_role_deleted_at_idx" ON "permission_role"("deleted_at");


CREATE TABLE IF NOT EXISTS "translations" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "translation" JSONB NOT NULL,
    UNIQUE ("name")

);

-- insert modules

INSERT INTO "module"(
    "id",
    "name",
    "app_type_id", 
    "created_by",
    "translation"
) VALUES (
    '5b690557-4461-4460-a079-56a7663566ab',
    'Products',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Mahusulotlar", "ru":"Товары", "en":"Products"}'
);

INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    'ca19a21a-664d-4ed1-8a49-eb2a1db2a38f',
    'Sale',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Sotuv", "ru":"Продажа", "en":"Sale"}'
);

INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    'c49304cb-db2c-4bb0-9845-8ef3c4b24910',
    'Finance',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Moliya", "ru":"Финансы", "en":"Finance"}'
);
INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    'beeb9c69-e709-4175-86f9-ec1963f16eea',
    'Management',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Boshqaruv", "ru":"Управление", "en":"Management"}'

);

INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    '77420ad4-2904-4c2f-b91f-b625d37e3c34',
    'Clients',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Klientlar", "ru":"Клиенты", "en":"Clients"}'
);
INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    '7521146a-e5bf-4b62-ad29-168a387c5a7c',
    'Settings',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Sozlamalar", "ru":"Настройки", "en":"Settings"}'
);
INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    '37956c85-afd1-4099-86b7-83944fdd1749',
    'Reports',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Hisobotlar", "ru":"Отчеты", "en":"Reports"}'

);
INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    'e92f2b76-7806-4dfc-9027-5bfc291692bd',
    'Kanban',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kanban", "ru":"Канбан", "en":"Kanban"}'
);
INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    'd547c9dd-3800-4c86-85f8-902ff2d44756',
    'Marketing',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Marketing", "ru":"Маркетинг", "en":"Marketing"}'
);
INSERT INTO "module" (
    "id",
    "name",
    "app_type_id",
    "created_by",
    "translation"
) VALUES (
    'a802b300-6e6a-42f2-99ce-bde7a206650d',
    'Billing',
    '1fe92aa8-2a61-4bf1-b907-182b497584ad',
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Hisob-kitob", "ru":"Биллинг", "en":"Billing"}'
);


-- insert products module sections 
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '62cdb9e3-38bf-4f5e-9dc1-977dd6a69276',
    'Catalog',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Katalog", "ru":"Каталог", "en":"Catalog"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '71967e9c-3a16-410c-82f0-608ff4bbe7cf',
    'Write off',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Write Off", "ru":"Cписание", "en":"Write Fff"}'
);
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '2dce3c0f-123c-4b5e-af52-e267437787f6',
    'Import',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Import", "ru":"Импорт", "en":"Import"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '23fc53a1-3a37-41c1-9ad0-e985067cb0ba',
    'Transfer',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Transfer", "ru":"Передача", "en":"Transfer"}'

);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '55d75b58-c151-4d7b-a619-eeb8065f381f',
    'Inventory',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Inventarizatsiya", "ru":"Инвентаризация", "en":"Inventory"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '54e0d4a1-2df2-42c8-90c3-efdf5cce74bd',
    'Orders',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Buyurtmalar", "ru":"Заказы", "en":"Orders"}'

); 

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '19d97f84-112f-4d8a-a3e6-787f3c218acf',
    'Repricing',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Repricing", "ru":"Переоценка", "en":"Repricing"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '8243efd3-58da-4c1a-a7c3-e85d3d1cc51d',
    'Label Print',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Label Print", "ru":"Печать Ценников", "en":"Yorliqlar Chiqarish"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '6175bc30-fc74-4546-aa28-98859e4935c1',
    'Suppliers',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Yetkazib beruvchilar", "ru":"Поставщики", "en":"Suppliers"}'

);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    'ce52c92f-2a8b-430c-8dd7-1093810ed5c5',
    'Categories',
    '5b690557-4461-4460-a079-56a7663566ab', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kategoriyalar", "ru":"Категории", "en":"Categories"}'

);


-- sales
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    'a4c673a9-0b0c-4689-9feb-bf954fec2e06',
    'New Sale',
    'ca19a21a-664d-4ed1-8a49-eb2a1db2a38f', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Yangi Sotuv", "ru":"Новая Продажа", "en":"New Sale"}'
);
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '8921a4ef-23b3-4608-93bb-54a16907a2e0',
    'All Sales',
    'ca19a21a-664d-4ed1-8a49-eb2a1db2a38f', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Barcha Sotuvlar", "ru":"Все продажи", "en":"All Sales"}'
); 

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '1202a5b0-dbc1-45f3-b774-84c5670eb7b1',
    'Shifts',
    'ca19a21a-664d-4ed1-8a49-eb2a1db2a38f', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Smenalar", "ru":"Смены", "en":"Shifts"}'
); 

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '445b8f2d-b13f-40d8-9946-489ca936926a',
    'Operations',
    'ca19a21a-664d-4ed1-8a49-eb2a1db2a38f', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Jarayonlar", "ru":"Операции", "en":"Operations"}'
); 

-- Management
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '93ebc9a0-9400-4864-af1c-0f0acaab943f',
    'Employees',
    'beeb9c69-e709-4175-86f9-ec1963f16eea', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Xodimlar", "ru":"Сотрудники", "en":"Employees"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '80816c70-ca54-4cd8-b8a7-4733799406f0',
    'Roles',
    'beeb9c69-e709-4175-86f9-ec1963f16eea', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Rollar", "ru":"Роли", "en":"Roles"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    'bf92111d-aa04-4421-8b31-ee3b6801d04e',
    'HR',
    'beeb9c69-e709-4175-86f9-ec1963f16eea', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kadrlar Bo''limi", "ru":"Отдел Кадров", "en":"HR"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    'c0fbf776-9466-43f7-9c4e-0f340cdd3f29',
    'Time cards',
    'beeb9c69-e709-4175-86f9-ec1963f16eea', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Voht Kartalari", "ru":"Временные Карты", "en":"Time Cards"}'
);

-- Clients
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '313db387-84c4-4365-99b0-a01ac85d12a0',
    'All clients',
    '77420ad4-2904-4c2f-b91f-b625d37e3c34', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Xamma Klientlar", "ru":"Все Клиенты", "en":"All Clients"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    'aff8591f-ec99-49ed-82de-c0a4ff1c058b',
    'Client debts',
    '77420ad4-2904-4c2f-b91f-b625d37e3c34', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Mijoz Qarzlari", "ru":"Задолжности Клиента", "en":"Client Debts"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '7beb33f0-7c28-420e-afe8-0c567c96f8fb',
    'Loyalty program',
    '77420ad4-2904-4c2f-b91f-b625d37e3c34', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Sodiqlik Dasturi", "ru":"Программа Лояльности", "en":"Loyalty Program"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '49085ba1-62a3-453c-a805-6f75d85c7475',
    'Client Groups',
    '77420ad4-2904-4c2f-b91f-b625d37e3c34', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Mijoz Guruhlari", "ru":"Группы Клиентов", "en":"Client Groups"}'
);

-- settings

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '63685ec6-72b2-41c4-a992-3e281397004e',
    'Company',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kompaniya", "ru":"Компания", "en":"Company"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '9541320f-f8bc-41ab-bc63-2dcc9397fb69',
    'Cashbox',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kassa", "ru":"Касса", "en":"Cashbox"}'
);
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '66a2ff50-52e1-499f-9c8a-5ffc1082e79f',
    'Store',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Do''konlar", "ru":"Магазины", "en":"Stores"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '5c8ae38d-98c9-4128-a157-677ad3b7c858',
    'Products',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Mahsulotlar", "ru":"Товары", "en":"Products"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    '2dfd0eb7-7ebf-47b3-abb7-fc7b348d74d8',
    'Application',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Ilova", "ru":"Приложение", "en":"Application"}'
);
INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES (
    'c4857445-4e5b-40a7-b51e-00c5c352edfe',
    'Payment',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"To''lov turlari", "ru":"Виды оплат", "en":"Payments"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    'Receipt',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kvitansiyalar", "ru":"Квитанции", "en":"Reciepts"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    '02662be0-c2c9-4291-ae0d-982d89e0f20b',
    'Labels',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Yorliq", "ru":"Ценник", "en":"Label"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    '4f7ac098-95a0-4ad8-9c5e-135772144002',
    'Vat',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Soliq", "ru":"Налог", "en":"Vat"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    '52ed72d4-815e-4d6d-b2ba-856535ae5534',
    'Currency',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Valyuta", "ru":"Валюта", "en":"Currency"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    'f1f1852e-34a8-4c4c-bf35-2d39314de31d',
    'Warehouse',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Ombor", "ru":"Склад", "en":"Warehouse"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    '32f5e61a-8f63-48ed-9c0f-9449ac5f74dc',
    'Profile',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Profil", "ru":"Профиль", "en":"Profile"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    'fad8853c-92bb-4794-8782-51e1388c5ae4',
    'Kanban',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Kanban", "ru":"Канбан", "en":"Kanban"}'
);

INSERT INTO "section" (
    "id",
    "name",
    "module_id",
    "created_by",
    "translation"
) VALUES(
    '3b0903bf-90cb-4970-a306-acc1c4239524',
    'Notification',
    '7521146a-e5bf-4b62-ad29-168a387c5a7c', 
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    '{"uz":"Bildirishnoma", "ru":"Уведомление", "en":"Notification"}'
);

-- Translations
INSERT INTO "translations" ("id", "name", "translation") VALUES
    ('56831a31-efcc-4402-be1b-811f0c54c8a9', 'required_field_error', '{"uz":"Maydonni to''ldirishingiz shart!", "en":"You should fill this field", "ru":"Вы должны заполнить это поле"}'),
    ('d8554ab4-d8f0-464e-bdb1-883c2733044e', 'min_6_chars_error', '{"uz":"Parol juda qisqa - kamida 6 ta belgidan iborat bo''lishi kerak", "en":"Password is too short - should be 6 chars minimum", "ru":"Пароль слишком короткий - должен быть минимум 6 символов"}'),
    ('fd844d4b-4f6f-45d1-b58e-adb2c9d239d5', 'hello', '{"uz":"Salom", "en":"Hello", "ru":"Здравствуйте"}'),
    ('98aa4fbf-13eb-4ad0-8961-3f0df2666baa', 'sign_in', '{"uz":"Tizimga kirish", "en":"Sign in", "ru":"Войти"}'),
    ('a3b264a4-4f89-43b3-9017-c9c19c9c1b62', 'sign_up', '{"uz":"Ro''yxatdan o''tish", "en":"Sign up", "ru":"Зарегистрироваться"}'),
    ('9aca5c64-e2c7-49d7-b318-7d382592ec45', 'new_to_invan', '{"uz":"INVAN da yangimisiz?", "en":"New to INVAN?", "ru":"Новый в INVAN?"}'),
    ('24026c29-24cd-49da-aa49-1cb8ee3589d8', 'phone_number', '{"uz":"Telefon raqam", "en":"Phone number", "ru":"Номер телефона"}'),
    ('0ffc98cb-6d0e-40cf-a313-e947c14fc780', 'password', '{"uz":"Parol", "en":"Password", "ru":"Пароль"}'),
    ('30f7019e-32b0-4ab0-b91d-ec614a4f112e', 'forgot_password', '{"uz":"Parolni unutdingizmi?", "en":"Forgot password", "ru":"Забыли пароль?"}'),
    ('433cc531-0b04-4ce2-a762-7152f458fcf7', 'enter_your_password', '{"uz":"Parolingizni kiriting", "en":"Enter your password", "ru":"Введите свой пароль"}'),
    ('a9bb3c73-4980-4c1e-a80c-ef152d2cb61f', 'enter_your_phone_number', '{"uz":"Telefon raqamingizni kiriting", "en":"Enter your phone number", "ru":"Введите свой номер телефона"}'),
    ('96789df0-f8d6-4508-b852-8056d4df0fc3', 'lets_create_your_account', '{"uz":"Keling, hisobingizni yarataylik", "en":"Let''s create your account.", "ru":"Давайте создадим вашу учетную запись."}'),
    ('0281e41d-a415-4a4a-89f1-b4dbe89e5e3a', 'signing_up_text', '{"uz":"INVAN da ro''yxatdan o''tish tez va bepul. Hech qanday majburiyat yoki uzoq muddatli shartnomalar yo''q.", "en":"Signing up for INVAN is fast and free. No commitments or long-term contracts.", "ru":"Зарегистрироваться в INVAN можно быстро и бесплатно. Никаких обязательств или долгосрочных контрактов."}'),
    ('5a031cbf-f0c5-431b-99cb-08ceb8f0eb46', 'continue', '{"uz":"Davom etish", "en":"Continue", "ru":"Продолжить"}'),
    ('5f6de14d-5578-4cd2-8efb-6c110743d927', 'already_have_an_account', '{"uz":"INVAN hisobingiz bormi?", "en":"Already have an INVAN account?", "ru":"У вас уже есть учетная запись INVAN?"}'),
    ('364583c1-5323-4c87-ac9e-220f170f167e', 'enter_code', '{"uz":"Kodni kiriting", "en":"Enter code", "ru":"Введите код"}'),
    ('361ad428-f583-4b38-b2ad-f98ea2431541', 'enter_code_that_you_received', '{"uz":"Raqamingizga kelgan kodni kiriting", "en":"Enter the code that you received on the number", "ru":"Введите код, который вы получили на номер"}'),
    ('55b18f34-9297-4048-80e5-564e661f428d', 'change_number', '{"uz":"Raqamni o''zgartirish", "en":"Change number", "ru":"Изменить номер"}'),
    ('ed2938c9-4bc1-4440-866b-f1ba7019cd74', 'didnt_receive_a_code', '{"uz":"Kodni olmadingizmi?", "en":"Didn''t receive a code?", "ru":"Не получили код?"}'),
    ('d4afcd65-2260-4095-8ce1-ae14de1057fd', 'we_can_resend_sms', '{"uz":"SMSni qayta yuborishimiz mumkin", "en":"We can resend SMS after", "ru":"Мы можем повторно отправить SMS"}'),
    ('d985d4fd-7f17-4d83-a4ac-d2a23e69732a', 'seconds', '{"uz":"sekundlar", "en":"seconds", "ru":"секунды"}'),
    ('21f87178-7887-4612-b5ee-bbbca1f4e72d', 'create_account', '{"uz":"Hisob yaratish", "en":"Create Account", "ru":"Создать аккаунт"}'),
    ('24e36f42-8d21-48cb-a240-c573d8b4c7f5', 'back', '{"uz":"Orqaga", "en":"Back", "ru":"Назад"}'),
    ('aa9c5349-6a22-4a0e-8d5c-a9cb199da732', 'learn_more', '{"uz":"Batafsil ma''lumot", "en":"Learn more", "ru":"Узнать больше"}'),
    ('a18f76bf-288e-47a7-afce-5e966965fce4', 'new_password', '{"uz":"Yangi parol", "en":"New Password", "ru":"Новый пароль"}'),
    ('65714bbf-e963-46d1-9da9-cd161524e7f6', 'confirm_password', '{"uz":"Parolni tasdiqlang", "en":"Confirm password", "ru":"Подтвердите пароль"}'),
    ('9a27ff3c-8a0d-4fa1-954c-84f95930f6a1', 'finish', '{"uz":"Tugatish", "en":"Finish", "ru":"Готово"}'),
    ('7b2158a6-7c49-45de-952f-e751fdf785ed', 'name', '{"uz":"Ism", "en":"Name", "ru":"Имя"}'),
    ('bd46aa65-e014-4ee5-89c6-f9cd4ccf3417', 'main', '{"uz":"Asosiy", "en":"Main", "ru":"Главный"}'),
    ('c52a90e2-6e6e-4d6c-89d5-4a80be3d925d', 'profile', '{"uz":"Profile", "en":"Profile", "ru":"Профиль"}'),
    ('c751134a-9226-4784-a3f5-c47e3f016014', 'reset', '{"uz":"Yangilash", "en":"Reset", "ru":"Сброс"}'),
    ('10687a6d-5d9d-4965-8a05-53af619898d6', 'save', '{"uz":"Saqlash", "en":"Save", "ru":"Сохранить"}'),
    ('3365cdbb-6aaf-49de-9ee3-7011c81e4cef', 'surname', '{"uz":"Familiya", "en":"Surname", "ru":"Фамилия"}'),
    ('7f0dec16-647e-4ec0-acb8-1230fbb53953', 'security', '{"uz":"Xavfsizlik", "en":"Security", "ru":"Безопасность"}'),
    ('3d9e098b-889f-432d-a3bb-5f825c939a12', 'old_password', '{"uz":"Eski parol", "en":"Old password", "ru":"Старый пароль"}'),
    ('b56835a8-9a92-4745-8dc4-19480f208dfe', 'enter_old_password', '{"uz":"Eski parolni kiriting", "en":"Enter old password", "ru":"Введите старый пароль"}'),
    ('cd80985e-2523-4edf-bd90-fa4fb37f4730', 'confirm_new_password', '{"uz":"Yangi parolni tasdiqlang", "en":"Confirm new password", "ru":"Подтвердите новый пароль"}'),
    ('3b545fcc-e503-4b2f-a3ce-5827db18927a', 'enter_new_password', '{"uz":"Yangi parolni kiriting", "en":"Enter new password", "ru":"Введите новый пароль"}'),
    ('30af6136-3f6c-44c0-a6b8-2bed837d9a96', 'interface', '{"uz":"Interfeys", "en":"Interface", "ru":"Интерфейс"}'),
    ('3a9f84fa-4dc3-462b-947c-cf8403c800ad', 'select_language', '{"uz":"Tilni tanlang", "en":"Select language", "ru":"Выберите язык"}'),
    ('4f5d5f14-27c3-4ca7-810e-caef4113006f', 'language', '{"uz":"Til", "en":"Language", "ru":"Язык"}'),
    ('f8555251-7532-4b24-9195-848b00642fa4', 'company', '{"uz":"Kompaniya", "en":"Company", "ru":"Компания"}'),
    ('95c5f71a-a6e4-480e-8b4b-b049e8fc4ee8', 'business_name', '{"uz":"Biznes nomi", "en":"Business name", "ru":"Наименование фирмы"}'),
    ('5188810e-e5dc-4bc5-95ba-127be1707a55', 'business_email', '{"uz":"Biznes pochta", "en":"Business email", "ru":"Рабочий адрес электронной почты"}'),
    ('c83bba81-5780-498f-9a76-d60100fc33ad', 'enter_business_email', '{"uz":"Biznes pochtani kiriting", "en":"Enter business email", "ru":"Введите рабочий адрес электронной почты"}'),
    ('f43b1a5a-940c-47d2-b9c1-d3197e2c7602', 'email', '{"uz":"Email", "en":"Email", "ru":"Email"}'),
    ('bd3a9e8e-34cb-4cea-bcf5-91feb7013b75', 'requisites', '{"uz":"Rekvizitlar", "en":"Requisites", "ru":"Реквизиты"}'),
    ('61d96dd1-45ba-4b37-959e-49e0ae88296a', 'enter_name', '{"uz":"Nomni kiriting", "en":"Enter name", "ru":"Введите имя"}'),
    ('e1e3cf98-6485-4f91-a7e4-8438f518eb95', 'legal_name', '{"uz":"Yuridik nomi", "en":"Legal name", "ru":"Юридическое название"}'),
    ('1f5ab09d-6808-4456-a79e-57aa2c96c501', 'legal_address', '{"uz":"Yuridik manzil", "en":"Legal address", "ru":"Юридический адрес"}'),
    ('4e408815-db88-4409-82f0-d98bed074ff1', 'country', '{"uz":"Davlat", "en":"Country", "ru":"Страна"}'),
    ('78f0c165-2b7c-45c7-88ac-870a0bf6a2d1', 'select_country', '{"uz":"Davlatni tanlang", "en":"Select country", "ru":"Выберите страну"}'),
    ('9b9cc41e-4d57-4910-add5-2eb60d50a30b', 'zip_code', '{"uz":"Zip code", "en":"Zip code", "ru":"Zip code"}'),
    ('7bda1275-c76e-4a5c-951f-2f78b2fb5ee3', 'enter_zip_code', '{"uz":"Zip kodni kiriting", "en":"Enter zip code", "ru":"Введите почтовый индекс"}'),
    ('6dd96247-d63c-4cb1-b84a-b282622cc2da', 'tin', '{"uz":"STIR (soliq to''lovchining identifikatsiya raqami)", "en":"TIN (Taxpayer identication number)", "ru":"ИНН (идентификационный номер налогоплательщика)"}'),
    ('dd5f94ee-860d-4981-b23f-d1d9dd9de758', 'enter_tin', '{"uz":"STIRni kiriting", "en":"Enter TIN", "ru":"Введите ИНН"}'),
    ('86838608-8f0e-4fe4-9390-df344a76d562', 'enter_ibt', '{"uz":"FA (filiallararo aylanma)", "en":"Enter IBT", "ru":"Введите МО"}'),
    ('a7b73300-490c-4378-94ee-ed2e9a46439e', 'store', '{"uz":"Do''kon", "en":"Store", "ru":"Магазин"}'),
    ('72996755-e91b-42aa-8729-568f369f6510', 'create', '{"uz":"Yaratish", "en":"Create", "ru":"Создать"}'),
    ('60d0b0a6-a1b3-4a19-ad0e-2ae5c43e891f', 'number_of_cashboxes', '{"uz":"Kassalar soni", "en":"Number of cashboxes", "ru":"Количество касс"}'),
    ('961d6428-1502-401e-a857-0a0b7a42fa91', 'create_store', '{"uz":"Magazin yaratish", "en":"Create store", "ru":"Создать магазин"}'),
    ('f66ad94d-0d8c-4816-b05c-6fde18b035f6', 'title', '{"uz":"Nomi", "en":"Title", "ru":"Название"}'),
    ('63d22859-649a-4b39-a666-5f3bd0db2102', 'enter_title', '{"uz":"Nomini kiriting", "en":"Enter title", "ru":"Введите название"}'),
    ('e5311870-629b-4a20-8e4b-e7847101d615', 'enter_m2', '{"uz":"m2 kiriting ", "en":"Enter m2", "ru":"Введите м2"}'),
    ('c758b8fc-f821-42f4-89a5-137bcc982bb8', 'enter_phone_number', '{"uz":"Telefon nomer kiriting", "en":"Enter phone number", "ru":"Введите номер телефона"}'),
    ('3d1d6c5a-ee2b-405c-87e2-941d2cb88f89', 'enter_address', '{"uz":"Manzil kiriting", "en":"Enter address", "ru":"Введите адрес"}'),
    ('4c4c5cc7-cfe8-49a6-925a-9653a23808f3', 'size', '{"uz":"O''lchami", "en":"Size", "ru":"Размер"}'),
    ('ea9c5d93-0f6d-466c-b6ab-468a821434f6', 'address', '{"uz":"Manzil", "en":"Address", "ru":"Адрес"}'),
    ('5e9d3fa1-158b-4ba0-9eea-69f51297d9ba', 'phone', '{"uz":"Telefon", "en":"Phone", "ru":"Телефон"}'),
    ('bdbe96c2-6af1-476a-8dcb-2fdb005fa867', 'information', '{"uz":"Ma''lumot", "en":"Information", "ru":"Информация"}'),
    ('6bc0e0b5-ad75-4020-a51a-0cb0c2adb0d9', 'enter_additional_information', '{"uz":"Qo''shimcha ma''lumot kiriting", "en":"Enter additional information", "ru":"Введите дополнительную информацию"}'),
    ('32339593-c7ba-4dba-9912-7616673e8a1a', 'cashbox', '{"uz":"Kassa", "en":"Cashbox", "ru":"Касса"}'),
    ('3a578edd-bcf1-4c54-9c38-d4b9451bbb87', 'create_cashbox', '{"uz":"Kassa yaratish", "en":"Create cashbox", "ru":"Создать кассу"}'),
    ('ad519291-b143-40ac-acab-f92748586e5d', 'select_store', '{"uz":"Magazinni tanlang", "en":"Select store", "ru":"Выберите магазин"}'),
    ('b0b348ba-bdb5-4511-916e-49812395a112', 'payment_types', '{"uz":"To''lov turlari", "en":"Payment types", "ru":"Типы платежей"}'),
    ('3f8f2b3d-b2e7-4c1e-8b33-ca0d668ead69', 'receipt_type', '{"uz":"Kvitansiya turi", "en":"Receipt type", "ru":"Тип чека"}'),
    ('d2faa347-9684-41ce-afbb-40c50d52ea28', 'integration_ofd', '{"uz":"Integratsiya(Ofd.uz)", "en":"Integration(Ofd.uz)", "ru":"Интеграция(Ofd.uz)"}'),
    ('48143e2d-3343-4a9a-910c-e42a658ef5ad', 'standart', '{"uz":"Standart", "en":"Standart", "ru":"Стандарт"}'),
    ('721b8e68-79ea-43ca-aebf-3cd34882c1b8', 'medium', '{"uz":"O''rta", "en":"Средний", "ru":"Средний"}'),
    ('2a9a060e-ac8b-4c0b-9bc5-a070d4732814', 'yes', '{"uz":"Ha", "en":"Yes", "ru":"Да"}'),
    ('0d7cb7fd-5246-482d-88ef-122a2786125c', 'No', '{"uz":"Yo''q", "en":"No", "ru":"Нет"}'),
    ('afb9bb7c-1e0b-40bb-94fe-3901ab4fdf28', 'products', '{"uz":"Mahsulotlar", "en":"Products", "ru":"Товары"}'),
    ('1aaac4e6-06dd-499b-be2b-0ec424ea095e', 'units', '{"uz":"Birliklar", "en":"Units", "ru":"Единицы"}'),
    ('bca18efd-6264-475c-8ad1-8f2dab563246', 'options', '{"uz":"Variantlar", "en":"Options", "ru":"Опции"}'),
    ('ee01f59d-88d9-4327-832a-e3da98ac65d7', 'custom_field', '{"uz":"Moslashtirilgan maydon", "en":"Custom field", "ru":"Пользовательское поле"}'),
    ('23628b99-39b7-4ce4-a3c8-dcdef3437ae5', 'add', '{"uz":"Qo''shish", "en":"Add", "ru":"Добавить"}'),
    ('620811f6-2f2b-4cc0-8a61-0df4e57c6fa4', 'invalid_phone_number', '{"uz":"No''tog''ri raqam", "en":"Invalid phone number", "ru":"Неверный номер телефона"}'),
    ('e19ad5a3-7dd8-44c8-a218-c0a2e5f31dfa', 'loading', '{"uz":"Yuklanmoqda", "en":"Loading", "ru":"Загрузка"}'),
    ('1b4006cd-3417-4b2d-be45-e29f77155ce7', 'incorrect_code', '{"uz":"Noto''g''ri kod", "en":"Incorrect code", "ru":"Неверный код"}'),
    ('710e5685-12d8-43d1-a050-4625eaa84683', 'already_registered', '{"uz":"Registratsiyadan otilgan", "en":"Already registered", "ru":"Уже зарегистрирован"}'),
    ('e60e9a3f-be35-451c-a689-12ede320d016', 'not_registered', '{"uz":"Registratsiyadan otilmagan", "en":"Not registered", "ru":"Не зарегистрирован"}'),
    ('87c4ac68-9e47-4bc8-8e9b-2449ac7f6322', 'receipt', '{"uz":"Retsept", "en":"Receipt", "ru":"Рецепт"}'),
    ('d277ab33-4739-463c-837e-4c6926f617b2', 'search', '{"uz":"Qidirish", "en":"Search", "ru":"Поиск"}'),
    ('7c86d95a-9bf0-43f2-b485-39d75840a50d', 'statistics', '{"uz":"Statistika", "en":"Statistics", "ru":"Статистика"}'),
    ('fdc6d4cf-17e8-4a00-86f9-073964852494', 'print_report', '{"uz":"Hisobotni chop etish", "en":"Print report", "ru":"Распечатать отчет"}'),
    ('080a2fc9-2b2e-4ac1-ace2-a34b711df1c7', 'sales_not_found', '{"uz":"Sotish topilmadi", "en":"Sales are not found", "ru":"Продажи не найдены"}'),
    ('176844a0-283c-4a92-ad07-c718031b0a31', 'change_password', '{"uz":"Parolni o''zgartirish", "en":"Change password", "ru":"Изменить пароль"}'),
    ('5949cee6-f5b7-4706-8032-28296551601d', 'change_interface', '{"uz":"Interfeysni o''zgartirish", "en":"Change interface", "ru":"Изменить интерфейс"}'),
    ('e1502873-4156-41cf-812c-fb3459b618a7', 'invalid_mxik_code', '{"uz":"MXIK kodi noto''g''ri", "en":"Invalid MXIK code", "ru":"Неверный MXIK код"}'),
    ('ee13845d-5ec6-4f1c-9cc4-b8e1c1fdedcf', 'remains_and_price', '{"uz":"Qoldiqlar va Narxi", "en":"Remains & Price", "ru":"Остатки и цена"}'),
    ('067fe4f9-b797-40f1-8347-07bc31359ee9', 'available', '{"uz":"Mavjud", "en":"Available", "ru":"Доступный"}'),
    ('04eebb69-c526-48e5-b7f0-0962892cffb5', 'supply_price', '{"uz":"Ta''minot narxi", "en":"Supply price", "ru":"Цена поставки"}'),
    ('3496d373-5fe9-435e-85b5-eb949ebfd900', 'retail_price', '{"uz":"Chakana savdo narxi", "en":"Retail price", "ru":"Розничная цена"}'),
    ('f4d3fe28-002b-4198-b8e5-8bd5ba85fffe', 'low_stock', '{"uz":"Kam zaxira", "en":"Low stock", "ru":"Низкий запас"}'),
    ('81dfcf9a-81bf-4a15-9b11-d9cc7a8103a2', 'delete', '{"uz":"O''chirish", "en":"Delete", "ru":"Удалить"}'),
    ('6f6370c8-71cb-4dc8-8031-c9ac0b589754', 'delete_the_category', '{"uz":"Kategoriyani o''chiring", "en":"Delete the category", "ru":"Удалить категорию"}'),
    ('46077307-b3ab-4d1d-a944-4cb64a419ac7', 'close', '{"uz":"Yopish", "en":"Close", "ru":"Закрыть"}'),
    ('e457d043-a456-4ffe-99e4-ec8918062d13', 'delete_the_category_text', '{"uz":"Haqiqatan ham turkumni oʻchirib tashlamoqchimisiz?", "en":"Are you sure you want to delete the category?", "ru":"Вы действительно хотите удалить категорию?"}'),
    ('51739222-2dcf-4714-8c7e-c884617880eb', 'category_subcategory', '{"uz":"Kategoriya, pastki kategoriya", "en":"Category, subcategory", "ru":"Категория, подкатегория"}'),
    ('88b3d318-df81-43f0-9e76-efcad2b85ac8', 'number_of_products', '{"uz":"Mahsulotlar soni", "en":"Number of products", "ru":"Количество продуктов"}'),
    ('cdbac36f-c497-448a-8da1-7cca96432bb7', 'actions', '{"uz":"Harakatlar", "en":"Actions", "ru":"Действии"}'),
    ('10fe08fe-70bd-4866-b25f-c56d08c12221', 'new_category', '{"uz":"Yangi kategoriya", "en":"New category", "ru":"Новая категория"}'),
    ('f603c5e8-db32-45cd-9d3b-4619fc19812b', 'enter_password', '{"uz":"Parolni kiriting", "en":"Enter password", "ru":"Введите пароль"}'),
    ('6a0c0aed-20f9-4529-8daf-fa11ff3e2ecd', 'first_name', '{"uz":"Ism", "en":"First name", "ru":"Имя"}'),
    ('568042ad-0df0-4737-a886-810901f27b61', 'last_name', '{"uz":"Familiya", "en":"Last name", "ru":"Фамилия"}'),
    ('2e13df1e-2d4c-4e23-81a8-155eef08b7c3', 'create_employee', '{"uz":"Xodim yaratish", "en":"Create employee", "ru":"Создать сотрудника"}'),
    ('23062e44-7614-423b-b699-1d1e4188e1b0', 'select_role', '{"uz":"Rolni tanlang", "en":"Select role", "ru":"Выберите роль"}'),
    ('6b5824f0-bc4a-4f53-8c52-21bb2870a2bd', 'role', '{"uz":"Rol", "en":"Role", "ru":"Роль"}'),
    ('b2a6bf28-53d1-4744-92b0-7187906f9f20', 'store_and_roles', '{"uz":"Do''kon va rol", "en":"Store and role", "ru":"Магазин и роль"}'),
    ('0c6d264d-b8c3-455c-91e4-d2691077fb6b', 'status', '{"uz":"Holat", "en":"Status", "ru":"Статус"}'),
    ('21774db0-3dcd-495a-a820-ae9f59ad4651', 'full_name', '{"uz":"To''liq ism sharif", "en":"Full name", "ru":"ФИО"}'),
    ('e52b571f-9a0a-4e59-9c2b-4c28cbcc25c5', 'id', '{"uz":"ID", "en":"ID", "ru":"Айди"}'),
    ('b50abdb7-b4bd-47c8-9bf5-3c94c4b7c67b', 'forgot_old_password', '{"uz":"Eski parolni unutdim", "en":"Forgot old password", "ru":"Забыл старый пароль"}'),
    ('ea5e4ab5-b3c7-418b-b3fa-e5a04c896d70', 'apply', '{"uz":"Murojaat qiling", "en":"Apply", "ru":"Подать заявление"}'),
    ('1199b648-ecdf-474b-80e4-c1a012cd01c5', 'theme', '{"uz":"Mavzu", "en":"Theme", "ru":"Тема"}'),
    ('25dd5e74-06bb-4a8d-873b-a53a383d3990', 'next', '{"uz":"Keyingisi", "en":"Next", "ru":"Следующий"}'),
    ('0c9ff42c-4703-4726-9523-d18c36cb0766', 'valid_email', '{"uz":"Yaroqli email boʻlishi kerak", "en":"Must be a valid email", "ru":"Должен быть действительный адрес электронной почты"}'),
    ('df4c3f14-dffe-4bbb-af7c-95fd39ece0a2', 'valid_number', '{"uz":"Raqam kiriting", "en":"Enter a number", "ru":"Введите число"}'),
    ('fb8053d8-8a4a-4aa6-aa82-07bf61db5aa0', 'add_unit', '{"uz":"Добавить единицу", "en":"Add unit", "ru":"Birlik qo''shing"}'),
    ('60467d69-7cfa-48a7-9968-affd398af956', 'select_unit', '{"uz":"Birlikni tanlang", "en":"Select unit", "ru":"Выберите единицу измерения"}'),
    ('d0ee1126-d2fb-46e3-a98c-8aecf11b45e5', 'abbreviation', '{"uz":"Qisqartirish", "en":"Abbreviation", "ru":"Сокращенное название"}'),
    ('54bff9ef-5eab-4880-b7ff-32aa3b48e1bc', 'accuracy', '{"uz":"Aniqlik", "en":"Accuracy", "ru":"Точность"}'),
    ('5fe4a305-4034-44ac-9ac1-5e325742b2ca', 'select_accuracy', '{"uz":"Aniqlikni tanlang", "en":"Select accuracy", "ru":"Выберите точность"}'),
    ('3e1d29e9-0ebf-4457-b535-612f866260a3', 'enter_text_to_display', '{"uz":"Ko''rsatish uchun matnni kiriting", "en":"Enter text to display", "ru":"Введите текст для отображения"}'),
    ('0b2e97c5-61bb-4322-87bc-31cccb8c5160', 'your_text', '{"uz":"Sizning matningiz", "en":"Your text", "ru":"Ваш текст"}'),
    ('5c1f3d08-6ecf-4faf-b4d4-be32053e8b35', 'subtotal', '{"uz":"Oraliq jami", "en":"Subtotal", "ru":"Промежуточный итог"}'),
    ('c09b0a6e-3d20-445c-980a-cacbf52cc9ff', 'total', '{"uz":"Jami", "en":"Total", "ru":"Общий"}'),
    ('daa90ecc-d6d2-45ca-beae-89b249675e49', 'discount', '{"uz":"Chegirma", "en":"Discount", "ru":"Скидка"}'),
    ('d8acb5f8-5c68-4378-badb-da310e6f6ada', 'information_block', '{"uz":"Axborot bloki", "en":"Information block", "ru":"Информационный блок"}'),
    ('3ca86c7d-6d22-45da-ad33-c40699a9818f', 'bottom_block', '{"uz":"Pastki blok", "en":"Bottom block", "ru":"Нижний блок"}');


-- trgiggers

-- defaults
CREATE OR REPLACE FUNCTION create_defaults()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
    "per" RECORD;
    "owner_role_id" UUID := uuid_generate_v4();
BEGIN
    INSERT INTO "role" ("id", "name", "company_id", "is_deletable", "created_by") VALUES  ("owner_role_id", 'Owner', NEW.id, FALSE, NEW.created_by);
    
    FOR "per" in SELECT "id" FROM "permission" WHERE "deleted_at" = 0
        LOOP
            INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("owner_role_id", "per"."id", NEW.created_by ); 
        END LOOP;

    INSERT INTO "user_role" ("role_id", "company_id", "user_id", "created_by") VALUES ("owner_role_id", NEW.id, NEW.created_by, New.created_by);
    INSERT INTO "company_user" ("id", "user_id", "company_id", "created_by", "status") VALUES ("owner_role_id", NEW.created_by, NEW.id, NEW.created_by, '9fb3ada6-a73b-4b81-9295-5c1605e54552');
	RETURN NEW;


END;
$$;


-- triggers
CREATE OR REPLACE TRIGGER create_defaults
    AFTER INSERT ON "company"
    FOR EACH ROW
    EXECUTE PROCEDURE create_defaults();


CREATE OR REPLACE FUNCTION update_employe()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN

  RETURN NEW;
END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER update_employe
    AFTER INSERT OR UPDATE ON "company_user"
    FOR EACH ROW
    EXECUTE PROCEDURE update_employe();

CREATE OR REPLACE FUNCTION create_company_user_external_id()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
  external_id INT := 0;

BEGIN
    SELECT
      COUNT(*) AS total
    FROM
      "company_user"
    WHERE
      "company_id" = NEW."company_id"
    INTO
      external_id;

    external_id := external_id + 1;

    NEW."external_id"=RIGHT(CONCAT('000000', external_id), 6);

    RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER create_company_user_external_id
    BEFORE INSERT ON "company_user"
    FOR EACH ROW
    EXECUTE PROCEDURE create_company_user_external_id();
