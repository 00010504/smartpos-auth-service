-- defaults
CREATE OR REPLACE FUNCTION create_cashier_defaults()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
    "cashier_role_id" UUID := uuid_generate_v4();
BEGIN
    INSERT INTO "role" ("id", "name", "company_id", "is_deletable", "created_by") VALUES  ("cashier_role_id", 'Cashier', NEW.id, FALSE, NEW.created_by);
    
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("cashier_role_id", '53c360f9-a10b-4f4b-8b52-670c3b8f87ca', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("cashier_role_id", '888fde2d-812c-4e5e-9e51-601e367eecbf', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("cashier_role_id", 'e84a7e7d-2c6e-441f-82b7-1eeaf3a361fe', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("cashier_role_id", 'f89ca40a-f2f9-4745-bebc-9cf293521c95', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("cashier_role_id", '6f79f732-777c-4c76-98d9-d519ebf7ec78', NEW.created_by );

	RETURN NEW;

END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER create_cashier_defaults
    AFTER INSERT ON "company"
    FOR EACH ROW
    EXECUTE PROCEDURE create_cashier_defaults();

-- defaults
CREATE OR REPLACE FUNCTION create_accountant_defaults()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
    "accountant_role_id" UUID := uuid_generate_v4();
BEGIN
    INSERT INTO "role" ("id", "name", "company_id", "is_deletable", "created_by") VALUES  ("accountant_role_id", 'Accountant', NEW.id, FALSE, NEW.created_by);
    
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '952d3177-1791-4b2a-9a13-e6e7899bf91c', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '2f154c89-a261-419d-8cfe-9b07783d2360', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'c2279c2b-3d2f-4e59-85fe-235924acec06', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '168d7e5d-d37a-4534-938e-8af1db9c22e2', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '01c7b0e9-0206-4fda-844d-ad098566326a', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '5eaff4af-09af-4ae9-90c1-669f333c538b', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'df0a0a3d-1173-4831-949f-f5961c6cde5c', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '1a9b575b-31f5-47b7-9056-8875889921cd', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'e7ffe144-757b-420e-b57d-4f04095b97e9', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '2b641aa6-26b1-4d5b-9fc4-f5538338aec2', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '53c360f9-a10b-4f4b-8b52-670c3b8f87ca', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '888fde2d-812c-4e5e-9e51-601e367eecbf', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'f89ca40a-f2f9-4745-bebc-9cf293521c95', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'e84a7e7d-2c6e-441f-82b7-1eeaf3a361fe', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '49f674fc-efbf-42fc-bb7c-2904c9c15a07', NEW.created_by );

	RETURN NEW;

END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER create_accountant_defaults
    AFTER INSERT ON "company"
    FOR EACH ROW
    EXECUTE PROCEDURE create_accountant_defaults();

-- defaults
CREATE OR REPLACE FUNCTION create_analyst_defaults()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
    "analyst_role_id" UUID := uuid_generate_v4();
BEGIN
    INSERT INTO "role" ("id", "name", "company_id", "is_deletable", "created_by") VALUES  ("analyst_role_id", 'Analyst', NEW.id, FALSE, NEW.created_by);
    
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '952d3177-1791-4b2a-9a13-e6e7899bf91c', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '2f154c89-a261-419d-8cfe-9b07783d2360', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'c2279c2b-3d2f-4e59-85fe-235924acec06', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '168d7e5d-d37a-4534-938e-8af1db9c22e2', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '01c7b0e9-0206-4fda-844d-ad098566326a', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '5eaff4af-09af-4ae9-90c1-669f333c538b', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'df0a0a3d-1173-4831-949f-f5961c6cde5c', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '1a9b575b-31f5-47b7-9056-8875889921cd', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'e7ffe144-757b-420e-b57d-4f04095b97e9', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '2b641aa6-26b1-4d5b-9fc4-f5538338aec2', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '53c360f9-a10b-4f4b-8b52-670c3b8f87ca', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '888fde2d-812c-4e5e-9e51-601e367eecbf', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'f89ca40a-f2f9-4745-bebc-9cf293521c95', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", 'e84a7e7d-2c6e-441f-82b7-1eeaf3a361fe', NEW.created_by );
        INSERT INTO "permission_role" ("role_id", "permission_id", "created_by") VALUES ("accountant_role_id", '49f674fc-efbf-42fc-bb7c-2904c9c15a07', NEW.created_by );

	RETURN NEW;

END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER create_analyst_defaults
    AFTER INSERT ON "company"
    FOR EACH ROW
    EXECUTE PROCEDURE create_analyst_defaults();