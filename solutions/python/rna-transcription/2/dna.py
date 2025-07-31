def to_rna(dna_strand):
    rna_strand = []
    dna_to_rna = {'G': 'C',
                  'C': 'G',
                  'A': 'U',
                  'T': 'A'}
    for dna_nucleotide in dna_strand:
        rna_strand.append(dna_to_rna[dna_nucleotide])

    return ''.join(rna_strand)