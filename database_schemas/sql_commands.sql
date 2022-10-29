
Regras de negocio:
	1. Cadastrar uma conta bancárias;
	2. Retornar extrato de uma conta bancária;
	3. Retornar extrato de transações de uma conta bancária;
	4. Realizar transferência de uma conta bancária para outra;
	5. Trazer violações de um usuário.

- Cadastrar uma conta bancárias:

	INSERT INTO users(name, email, cpf) VALUES
		('ciclano', 'ciclano@gmail.com', '321');

	INSERT INTO accounts(id, active, available_limit) VALUES
		(1, true, 10000);
	
	INSERT INTO transactions(amount, merchant, time) VALUES
		(10, 'Mercado Livre', 'Foda-se'),
		(5, 'Pichau', 'XPTO');
	
	INSERT INTO clients_transactions(transaction_id, user_id) VALUES 
		(1, 1),
		(2, 1);

	INSERT INTO violations(user_id, violation) VALUES
		(1, 'Nao foi possivel realizar a transacao, voces ja fez uma em menos de 1 minuto'),
		(1, 'Voce nao tem limite');

- Retorna as transações de um usuario:

	SELECT t.amount, t.merchant, t.time FROM transactions t JOIN clients_transactions ct
		ON ct.transaction_id = t.id WHERE ct.user_id = 1;

- Retorna extrato da conta de um cliente:

	SELECT a.available_limit, u.name, u.email from accounts a JOIN users u
		ON a.id = u.id WHERE u.id = 1;

- Transferencia:

	# Subtração
	UPDATE accounts SET available_limit =
	CASE
		WHEN available_limit > 0 THEN available_limit - 10
	ELSE 0
	END
	WHERE id = 1;

	UPDATE accounts SET available_limit =
	CASE
		WHEN available_limit > 0 THEN available_limit - 5
	ELSE 0
	END
	WHERE id = 1;

	# Adição
	UPDATE accounts SET available_limit =
	CASE
		WHEN available_limit > 0 THEN available_limit + ?
	ELSE 0
	END
	WHERE id = ?

# Violations

	SELECT v.violation from violations v JOIN users u
		ON v.user_id = u.id WHERE u.id = 1;
