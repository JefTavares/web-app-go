Create TABLE produtos (
	id SERIAL PRIMARY KEY,
	nome VARCHAR,
	descricao VARCHAR ,
	preco DECIMAL,
	quantidade  integer
);

INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES 
('Camiseta', 'Preta', 19, 10), 
('Fone', 'Muito bom', 99, 5);