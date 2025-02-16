ALTER TABLE
    user_invitations
ADD COLUMN IF NOT EXISTS
    expiration_date timestamp(0) with time zone NOT NULL;