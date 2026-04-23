const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

function isValidUsername(username, validate) {
  if (!validate) {
    return true;
  }
  
  if (username.length < 6 || username.length > 30) {
    return false;
  }
  
  if (username.includes('..')) {
    return false;
  }
  
  if (username.startsWith('.') || username.endsWith('.')) {
    return false;
  }
  
  const forbiddenChars = ['&', '=', '+', '<', '>', ',', '_', "'", '-'];
  for (const char of forbiddenChars) {
    if (username.includes(char)) {
      return false;
    }
  }
  
  for (let i = 0; i < username.length; i++) {
    const char = username[i];
    if (!((char >= 'a' && char <= 'z') || 
          (char >= 'A' && char <= 'Z') || 
          (char >= '0' && char <= '9') || 
          char === '.')) {
      return false;
    }
  }
  
  return true;
}

function normalizeEmail(email, validate) {
  const [local, domain] = email.split('@');
  
  if (!isValidUsername(local, validate)) {
    return null;
  }
  
  let normalizedLocal = '';
  let ignore = false;
  
  for (let i = 0; i < local.length; i++) {
    const char = local[i];
    
    if (char === '+') {
      ignore = true;
      continue;
    }
    
    if (ignore) {
      continue;
    }
    
    if (char !== '.') {
      normalizedLocal += char;
    }
  }
  
  return normalizedLocal + '@' + domain;
}

function countUniqueEmails(emails, validate) {
  const uniqueSet = new Set();
  
  for (const email of emails) {
    const normalized = normalizeEmail(email, validate);
    if (normalized !== null) {
      uniqueSet.add(normalized);
    }
  }
  
  return uniqueSet.size;
}

function main() {
  const validate = process.argv.includes('--validate');
  
  rl.on('line', (line) => {
    if (line.trim() === '') {
      return;
    }
    
    const emails = line.split(',').map(e => e.trim()).filter(e => e);
    const result = countUniqueEmails(emails, validate);
    console.log(result);
  });
  
  rl.on('close', () => {
    process.exit(0);
  });
}

if (require.main === module) {
  main();
}