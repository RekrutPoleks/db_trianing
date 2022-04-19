DROP TABLE IF EXISTS users, tasks,labels,tasks_labels;

CREATE TABLE users(
                      id SERIAL PRIMARY KEY,
                      name text NOT NULL
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       opened BIGINT DEFAULT extract(epoch from now()),
                       closed BIGINT DEFAULT 0,
                       author_id INTEGER NOT NULL REFERENCES users(id)
                           ON DELETE CASCADE
                           ON UPDATE CASCADE,
                       assigned_id INTEGER DEFAULT 0 REFERENCES users(id)
                           ON UPDATE CASCADE,
                       tittle text NOT NULL,
                       content text
);

CREATE TABLE labels(
                       id SERIAL PRIMARY KEY,
                       name text
);

CREATE TABLE tasks_labels(
                             task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE,
                             label_id INTEGER REFERENCES labels(id) ON DELETE CASCADE
);

--INSERT
TRUNCATE users,labels,tasks,tasks_labels;

INSERT INTO users(id, "name") VALUES (0,'default');
INSERT INTO users("name") VALUES ('Auth1');
INSERT INTO users("name") VALUES ('Auth2');
INSERT INTO users("name") VALUES ('Auth3');
INSERT INTO users("name") VALUES ('Auth4');

INSERT INTO labels (name) VALUES ('Lbl1');
INSERT INTO labels (name) VALUES ('Lbl2');
INSERT INTO labels (name) VALUES ('Lbl3');
INSERT INTO labels (name) VALUES ('Lbl4');
INSERT INTO labels (name) VALUES ('Lbl5');
INSERT INTO labels (name) VALUES ('Lbl6');
INSERT INTO labels (name) VALUES ('Lbl7');

INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (0,1,'ttl_t1', 't1');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (1,1,'ttl_t2', 't2');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (2,1,'ttl_t3', 't3');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (3,1,'ttl_t4', 't4');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (4,1,'ttl_t5', 't5');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (0,1,'ttl_t6', 't6');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (1,1,'ttl_t7', 't7');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (2,1,'ttl_t8', 't8');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (3,1,'ttl_t9', 't9');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (4,1,'ttl_t10', 't10');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (0,1,'ttl_t11', 't11');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (1,1,'ttl_t12', 't12');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (2,1,'ttl_t13', 't13');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (3,1,'ttl_t14', 't14');
INSERT INTO tasks (author_id, assigned_id,tittle,content) VALUES (4,1,'ttl_t15', 't15');



INSERT INTO tasks_labels(task_id, label_id) VALUES(1,1);
INSERT INTO tasks_labels(task_id, label_id) VALUES(2,2);
INSERT INTO tasks_labels(task_id, label_id) VALUES(3,3);
INSERT INTO tasks_labels(task_id, label_id) VALUES(4,4);
INSERT INTO tasks_labels(task_id, label_id) VALUES(5,5);
INSERT INTO tasks_labels(task_id, label_id) VALUES(6,6);
INSERT INTO tasks_labels(task_id, label_id) VALUES(7,7);
INSERT INTO tasks_labels(task_id, label_id) VALUES(8,1);
INSERT INTO tasks_labels(task_id, label_id) VALUES(9,2);
INSERT INTO tasks_labels(task_id, label_id) VALUES(10,3);
INSERT INTO tasks_labels(task_id, label_id) VALUES(11,4);
INSERT INTO tasks_labels(task_id, label_id) VALUES(12,5);
INSERT INTO tasks_labels(task_id, label_id) VALUES(13,6);
INSERT INTO tasks_labels(task_id, label_id) VALUES(14,7);
INSERT INTO tasks_labels(task_id, label_id) VALUES(15,1);
