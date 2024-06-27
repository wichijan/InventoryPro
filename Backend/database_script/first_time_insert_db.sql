use InventoryProDB;

Insert into
    roles(id, role_name)
values
    (
        "b01ccc07-693f-42de-b7f5-f2b5307bbcc1",
        'Admin'
    );

Insert into
    user_types(id, type_name)
values
    (
        "7209efef-5a3f-4fab-8a47-7a157c2df829",
        'Teacher'
    );

Insert into
    users(
        id,
        first_name,
        last_name,
        username,
        email,
        phone_number,
        user_type_id,
        is_active,
        registration_accepted
    )
values
    (
        "d1592a60-1538-4d6a-b3fd-60193622a854",
        'Admin',
        'Admin',
        'Admin',
        'sebastian.kiebert.dhbw@gmail.com',
        "+0161838382920",
        "7209efef-5a3f-4fab-8a47-7a157c2df829",
        0,
        1
    );

Insert into
    user_roles(user_id, role_id)
values
    (
        "d1592a60-1538-4d6a-b3fd-60193622a854",
        "b01ccc07-693f-42de-b7f5-f2b5307bbcc1"
    );

Insert into
    registration_codes(user_id, code)
values
    ("d1592a60-1538-4d6a-b3fd-60193622a854", "187");