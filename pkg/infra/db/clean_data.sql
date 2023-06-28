-- Delete a specific row from the campaign table
DELETE FROM campaign
WHERE campaign_id = 1;

-- Delete all rows from the campaign_image table
DELETE FROM campaign_image;

-- Delete rows from the transaction table based on a condition
DELETE FROM public.transaction
WHERE status = 'completed';

-- Delete all rows from the user table
DELETE FROM "user";



DELETE FROM campaign_image;
DELETE FROM transaction;
DELETE FROM campaign;
DELETE FROM "user";

