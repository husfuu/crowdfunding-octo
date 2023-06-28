-- Insert data into the campaign table
INSERT INTO campaign (campaign_id, user_id, name, "desc", short_desc, perks, backer_count, goal_amount, current_amount, slug)
VALUES
  (1, 1, 'Campaign 1', 'Campaign 1 description', 'Short description 1', 'Perks 1', 10, 1000, 500, 'campaign-1'),
  (2, 2, 'Campaign 2', 'Campaign 2 description', 'Short description 2', 'Perks 2', 5, 2000, 1500, 'campaign-2'),
  (3, 1, 'Campaign 3', 'Campaign 3 description', 'Short description 3', 'Perks 3', 20, 5000, 3000, 'campaign-3'),
  (4, 3, 'Campaign 4', 'Campaign 4 description', 'Short description 4', 'Perks 4', 15, 3000, 2000, 'campaign-4');


-- Insert data into the campaign_image table
INSERT INTO campaign_image (campaign_image_id, campaign_id, file_name)
VALUES
  (1, 1, 'image1.jpg'),
  (2, 2, 'image2.jpg'),
  (3, 3, 'image3.jpg'),
  (4, 4, 'image4.jpg');


-- Insert data into the transaction table
INSERT INTO public.transaction (trx_id, campaign_id, user_id, amount, status, code, payment_url)
VALUES
  (1, 1, 1, 100, 'completed', '12345', 'https://example.com/payment1'),
  (2, 2, 3, 200, 'completed', '56789', 'https://example.com/payment2'),
  (3, 3, 2, 300, 'pending', 'ABCDE', 'https://example.com/payment3'),
  (4, 4, 1, 150, 'completed', 'FGHIJ', 'https://example.com/payment4');


-- Insert data into the user table
INSERT INTO "user" (user_id, name, occupation, email, password_hash, avatar_file_name, role)
VALUES
  (1, 'John Doe', 'Software Developer', 'john.doe@example.com', 'password123', 'avatar1.jpg', 'user'),
  (2, 'Jane Smith', 'Marketing Specialist', 'jane.smith@example.com', 'pass456', 'avatar2.jpg', 'user'),
  (3, 'Mike Johnson', 'Graphic Designer', 'mike.johnson@example.com', 'pwd789', 'avatar3.jpg', 'user'),
  (4, 'Emily Davis', 'Project Manager', 'emily.davis@example.com', 'p@ssword', 'avatar4.jpg', 'user');
