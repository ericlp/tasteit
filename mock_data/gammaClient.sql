-- Make sure the mock client does not exist
BEGIN;

DELETE FROM it_user_approval
WHERE itclient_id='21347b86-9f58-4fec-94f2-ce412cd95794';

DELETE FROM itclient
WHERE id='21347b86-9f58-4fec-94f2-ce412cd95794';

DELETE FROM internal_text
WHERE id='8c5ad5bf-4f41-4a0d-81f5-0280953ebb58';

COMMIT;


-- Insert the client
BEGIN;

INSERT INTO internal_text (id, sv, en) 
VALUES ('8c5ad5bf-4f41-4a0d-81f5-0280953ebb58',	'Klient för att mocka name',	'Client for mocking name');

INSERT INTO itclient (id, client_id, client_secret, web_server_redirect_uri, access_token_validity, refresh_token_validity, auto_approve, name, description, created_at, last_modified_at)
VALUES ('21347b86-9f58-4fec-94f2-ce412cd95794',	'client_id', '{noop}secret', 'http://localhost:3000/api/auth/account/callback',	3600,	500000000,	'0',	'tasteit',	'8c5ad5bf-4f41-4a0d-81f5-0280953ebb58',	'2022-08-22 20:17:58.300228', '2022-08-22 20:17:58.30023');

COMMIT;