use InventoryProDB;

Insert into
    roles(id, role_name)
values
    (
        "b01ccc07-693f-42de-b7f5-f2b5307bbcc1",
        'Admin'
    ),
    (UUID(), 'smallAdmin'),
    (UUID(), 'Teacher'),
    (UUID(), 'Jumper'),
    (UUID(), 'Student'),
    (UUID(), 'Intern');

Insert into
    item_status(id, status_name)
values
    (
        "740d9d99-7ac9-4f44-a67f-cdc2b7dad0a1",
        'Available'
    ),
    (UUID(), 'Reserved'),
    (UUID(), 'Borrowed'),
    (UUID(), 'Lost');

Insert into
    subjects(id, name, description)
values
    (
        UUID(),
        'Math',
        'Mathematics is the study of numbers, quantities, and shapes.'
    ),
    (
        UUID(),
        'Science',
        'Science is the study of the nature and behavior of natural things and the knowledge that we obtain about them.'
    ),
    (
        UUID(),
        'History',
        'History is the study of past events, particularly in human affairs.'
    ),
    (
        UUID(),
        'English',
        'English is the study of the English language and literature.'
    ),
    (
        UUID(),
        'Art',
        'Art is the expression or application of human creative skill and imagination.'
    ),
    (
        "e1234a08-cd29-4f8a-9d8e-2716d6d2b546",
        'Reading',
        'Reading is the action or skill of reading written or printed matter silently or aloud.'
    );

Insert into
    warehouses(id, name, description)
values
    (
        "150ac215-87c3-40d5-8c09-11fbefd58ed1",
        'Main Warehouse',
        'Main warehouse for storing all items.'
    ),
    (
        "0aa6ded5-0529-487c-a82e-7fb12b66aa0d",
        'Warehouse 1',
        'Warehouse 1 for storing all items.'
    ),
    (
        "0aa43435-3e67-4b34-9b4a-632dd017eb84",
        'Warehouse 2',
        'Warehouse 2 for storing all items.'
    );

Insert into
    rooms(id, name, warehouse_id)
values
    (
        "30ed1352-4d0b-4508-8470-2cf7a4362414",
        'Hall',
        "150ac215-87c3-40d5-8c09-11fbefd58ed1"
    ),
    (
        "62144dad-ea04-4591-921c-a37c85ed9226",
        'Room 1',
        "150ac215-87c3-40d5-8c09-11fbefd58ed1"
    );

Insert into
    shelve_types(id, type_name)
values
    (
        "a2ef8893-67c4-4703-b29c-b0274eafdc69",
        'Small'
    ),
    (UUID(), 'Medium'),
    (UUID(), 'Large'),
    (
        "17ee7aea-857e-4f5b-b1cb-4d3ce5b49517",
        'SchnellRegal'
    );

Insert into
    shelves(id, shelve_type_id, room_id)
values
    (
        "0a75123a-9736-4c60-a10c-16e76aced3d2",
        "a2ef8893-67c4-4703-b29c-b0274eafdc69",
        /* Small Shelve */
        "62144dad-ea04-4591-921c-a37c85ed9226"
        /* Room 1 */
    ),
    (
        UUID(),
        "17ee7aea-857e-4f5b-b1cb-4d3ce5b49517",
        /* SchnellRegal */
        "30ed1352-4d0b-4508-8470-2cf7a4362414"
        /* Hall */
    );

Insert into
    items(
        id,
        name,
        description,
        class_one,
        class_two,
        class_three,
        class_four,
        damaged,
        damaged_description,
        quantity,
        status_id
    )
values
    (
        "d2b6b0f5-a9b9-4ab5-9644-75407c9e9dfb",
        'Harry Potter 1',
        'Harry Potter and the Philosopher''s Stone is a fantasy novel written by British author J. K. Rowling.',
        1,
        0,
        0,
        0,
        0,
        /* false */
        Null,
        1,
        "740d9d99-7ac9-4f44-a67f-cdc2b7dad0a1"
    ),
    (
        "2cbdabf2-f8f4-4aab-8d2e-59ef464abf6c",
        'Pen',
        'A pen is a writing instrument used to apply ink to a surface, usually paper, for writing or drawing.',
        0,
        1,
        0,
        0,
        0,
        NULL,
        20,
        NULL
    ),
    (
        "7791a294-94de-4cbf-9243-3f7210664f92",
        'Pencil',
        'A pencil is an implement for writing or drawing, constructed of a narrow, solid pigment core in a protective casing that prevents the core from being broken or marking the user’s hand.',
        0,
        0,
        1,
        0,
        0,
        NULL,
        30,
        NULL
    ),
    (
        "9f78d8b6-8605-4f58-bdee-b64bc8fde0cb",
        'Eraser',
        'An eraser, also known as a rubber, is an article of stationery that is used for removing marks from paper or skin.',
        0,
        0,
        0,
        1,
        0,
        NULL,
        40,
        NULL
    );

Insert into
    items_in_shelve(item_id, shelve_id)
values
    (
        "9f78d8b6-8605-4f58-bdee-b64bc8fde0cb",
        "0a75123a-9736-4c60-a10c-16e76aced3d2"
        /* Small Shelve */
    ),
    (
        "7791a294-94de-4cbf-9243-3f7210664f92",
        "0a75123a-9736-4c60-a10c-16e76aced3d2"
        /* Small Shelve */
    ),
    (
        "d2b6b0f5-a9b9-4ab5-9644-75407c9e9dfb",
        "0a75123a-9736-4c60-a10c-16e76aced3d2"
        /* Small Shelve */
    ),
    (
        "2cbdabf2-f8f4-4aab-8d2e-59ef464abf6c",
        "0a75123a-9736-4c60-a10c-16e76aced3d2"
        /* Small Shelve */
    );

Insert into
    item_subjects(item_id, subject_id)
values
    (
        "d2b6b0f5-a9b9-4ab5-9644-75407c9e9dfb",
        "e1234a08-cd29-4f8a-9d8e-2716d6d2b546"
    );

Insert into
    keywords
values
    ("75a570b7-ca2a-4e36-82c3-bab80b65bceb", "light");

Insert into
    keywords_for_items (keyword_id, item_id)
values
    (
        "75a570b7-ca2a-4e36-82c3-bab80b65bceb",
        "d2b6b0f5-a9b9-4ab5-9644-75407c9e9dfb"
    );

Insert into
    user_types(id, type_name)
values
    (UUID(), 'Teacher'),
    (
        "7209efef-5a3f-4fab-8a47-7a157c2df829",
        'AdminAdmin'
    );

Insert into
    users(
        id,
        first_name,
        last_name,
        username,
        email,
        password,
        user_type_id
    )
values
    (
        "d1592a60-1538-4d6a-b3fd-60193622a854",
        'Admin',
        'Admin',
        'Admin',
        'admin@admin.com',
        null,
        "7209efef-5a3f-4fab-8a47-7a157c2df829"
    );

Insert into
    user_roles(user_id, role_id)
values
    (
        "d1592a60-1538-4d6a-b3fd-60193622a854",
        "b01ccc07-693f-42de-b7f5-f2b5307bbcc1"
    );