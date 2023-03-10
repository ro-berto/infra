-- Copyright 2016 The Chromium Authors
-- Use of this source code is governed by a BSD-style license that can be
-- found in the LICENSE file.


DROP PROCEDURE IF EXISTS CopyCommentToCommentContent;

delimiter //

CREATE PROCEDURE CopyCommentToCommentContent(
    IN in_start INT, IN in_stop INT, IN in_step INT)
BEGIN
  comment_loop: LOOP
    IF in_start >= in_stop THEN
      LEAVE comment_loop;
    END IF;

    SELECT in_start AS StartingAt;
    SELECT count(*)
    FROM Comment
    WHERE Comment.id >= in_start
    AND Comment.id < in_start + in_step;

    INSERT INTO CommentContent (comment_id, content, inbound_message)
    SELECT id, content, inbound_message
    FROM Comment
    WHERE Comment.id >= in_start
    AND Comment.id < in_start + in_step;

    SET in_start = in_start + in_step;

  END LOOP;

END;


//


delimiter ;


-- Copy and paste these individually and verify that the site is still responsive.
-- the first one, takes about 30 sec, then 4-7 minutes for each of the rest.
-- CALL CopyCommentToCommentContent(           0, 13 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(13 * 1000000, 16 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(16 * 1000000, 17 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(17 * 1000000, 18 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(18 * 1000000, 19 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(19 * 1000000, 20 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(20 * 1000000, 21 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(21 * 1000000, 22 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(22 * 1000000, 23 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(23 * 1000000, 24 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(24 * 1000000, 25 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(25 * 1000000, 26 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(26 * 1000000, 27 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(27 * 1000000, 28 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(28 * 1000000, 29 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(29 * 1000000, 30 * 1000000, 10000);
-- CALL CopyCommentToCommentContent(30 * 1000000, 40 * 1000000, 10000);
