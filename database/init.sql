CREATE DATABASE IF NOT EXISTS todo_list_db;
USE todo_list_db;

CREATE TABLE IF NOT EXISTS `tasks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(45) NOT NULL,
  `description` varchar(255) NOT NULL,
  `completed` tinyint NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Insert test data
INSERT INTO `tasks` (`title`, `description`, `completed`) VALUES
('Buy groceries', 'Milk, Bread, Eggs, and Vegetables', 0),
('Play some League of Legends', 'Ranked games with friends', 1),
('Call mom', 'Check in on her health', 1),
('Clean the house', 'Living room and kitchen', 0),
('Read a book', 'Finish a reading a book', 1),
