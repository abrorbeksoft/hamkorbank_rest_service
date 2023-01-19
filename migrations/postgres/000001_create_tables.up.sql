CREATE TABLE IF NOT EXISTS "phones"
(
    "id"         UUID PRIMARY KEY default gen_random_uuid(),
    "phone"      varchar                                    not null,
    "created_at" TIMESTAMP        DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP        DEFAULT CURRENT_TIMESTAMP NOT NULL
);

insert into phones (phone)
values ('+998931000000'),
       ('+998931000001'),
       ('+998931000002'),
       ('+998931000003'),
       ('+998931000004'),
       ('+998931000005'),
       ('+998931000006'),
       ('+998931000007'),
       ('+998931000008'),
       ('+998931000009'),
       ('+998931000010'),
       ('+998931000011'),
       ('+998931000012'),
       ('+998931000013'),
       ('+998931000014'),
       ('+998931000015'),
       ('+998931000016'),
       ('+998931000017'),
       ('+998931000018'),
       ('+998931000019'),
       ('+998931000020');